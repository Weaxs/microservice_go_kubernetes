package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
	"github.com/asaskevich/govalidator"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sync"
	"time"
)

const (
	DefaultProductFrozenExpires = 2 * 60 * 1000
)

var (
	warehouseClient, _ = NewWarehouseClient()
	paymentService     = &PaymentService{paymentRepo: &PaymentRepo{}}
	lockMap            = make(map[string]*sync.Mutex)
	settlementCache    = cache.New(4*60*time.Second, 8*60*time.Second)
)

type PaymentService struct {
	paymentRepo *PaymentRepo
	walletRepo  *WalletRepo
}

// ExecuteBySettlement 根据结算清单的内容执行，生成对应的支付单
func (s *PaymentService) ExecuteBySettlement(ctx context.Context, bill *domain.Settlement) (payment *domain.Payment, err error) {
	if _, err = govalidator.ValidateStruct(bill); err != nil {
		return
	}
	if len(bill.Items) < 1 {
		err = errors.New("结算单中缺少商品清单")
		return
	}

	err = s.replenishProductInformation(ctx, bill)
	if err != nil {
		return
	}
	payment, err = s.producePayment(ctx, bill)
	if err != nil {
		return
	}
	go func() {
		_ = s.setupAutoThawedTrigger(ctx, payment)
	}()
	return
}

func (s *PaymentService) CancelPayment(ctx context.Context, payId string) (err error) {
	l, ok := lockMap[payId]
	if !ok {
		l = &sync.Mutex{}
		lockMap[payId] = l
	}
	l.Lock()

	po, err := s.paymentRepo.GetByPayId(ctx, payId)
	if err != nil {
		return err
	}
	payment := convertPo(po)
	if payment.PayState == domain.PaymentState_WAITING {
		payment.PayState = domain.PaymentState_CANCEL
		err = s.paymentRepo.Update(ctx, convertPayment(payment))
		if err != nil {
			return err
		}
		err = s.accomplishSettlement(ctx, domain.PaymentState_CANCEL, payId)
		if err != nil {
			return err
		}
		klog.CtxInfof(ctx, "编号为%s的支付单已被取消", payId)
	} else {
		err = errors.New(fmt.Sprintf("当前订单不允许取消，当前状态为：%s",
			domain.PaymentState_name[int32(payment.PayState)]))
		return
	}
	l.Unlock()
	delete(lockMap, payId)

	// 取消成功清除缓存
	settlementCache.Delete(payId)
	return
}

// AccomplishPayment 完成支付 立即取消解冻定时器，执行扣减库存和资金
func (s *PaymentService) AccomplishPayment(ctx context.Context, payId string, accountId int64) (err error) {
	l, ok := lockMap[payId]
	if !ok {
		l = &sync.Mutex{}
		lockMap[payId] = l
	}
	l.Lock()

	po, err := s.paymentRepo.GetByPayId(ctx, payId)
	if err != nil {
		return err
	}
	payment := convertPo(po)
	if payment.PayState == domain.PaymentState_WAITING {
		payment.PayState = domain.PaymentState_PAYED
		err = s.paymentRepo.Update(ctx, convertPayment(payment))
		if err != nil {
			return err
		}
		err = s.accomplishSettlement(ctx, domain.PaymentState_PAYED, payId)
		if err != nil {
			return err
		}
		klog.CtxInfof(ctx, "编号为%s的支付单已处理完成，等待支付", payId)
	} else {
		err = errors.New(fmt.Sprintf("当前订单不允许支付，当前状态为：%s",
			domain.PaymentState_name[int32(payment.PayState)]))
		return
	}
	l.Unlock()
	delete(lockMap, payId)

	wallet, err := s.walletRepo.FindByAccountId(ctx, accountId)
	if err != nil {
		return err
	}
	if wallet.Money > payment.TotalPrice {
		wallet.Money = wallet.Money - payment.TotalPrice
		_ = s.walletRepo.Save(ctx, wallet)
		klog.CtxInfof(ctx, "支付成功。用户余额：%v，本次消费：%v", wallet.Money, payment.TotalPrice)
	} else {
		err = errors.New("用户余额不足以支付，请先充值")
	}

	if err != nil {
		// 扣款失败，恢复库存
		_ = s.rollbackSettlement(ctx, domain.PaymentState_PAYED, payId)
		return
	}

	// 支付成功清除缓存
	settlementCache.Delete(payId)
	return
}

func (s *PaymentService) replenishProductInformation(ctx context.Context, bill *domain.Settlement) (err error) {
	products, err := warehouseClient.GetProducts(ctx)
	if err != nil {
		return
	}

	bill.ProductMap = make(map[int64]*domain.Product)
	for _, product := range products {
		bill.ProductMap[product.Id] = product
	}
	return
}

func (s *PaymentService) producePayment(ctx context.Context, bill *domain.Settlement) (payment *domain.Payment, err error) {
	total := float64(0)
	for _, item := range bill.GetItems() {
		product, ok := bill.ProductMap[item.ProductId]
		if !ok {
			continue
		}
		warehouseClient.FrozenStockpile(ctx, item.ProductId, item.Amount)
		total += product.Price * float64(item.Amount)
	}
	//  12元固定运费
	total += 12

	payId := uuid.New().String()
	paymentLink := fmt.Sprintf("/pay/modify/%s?state=PAYED", payId)
	if ctx.Value("account") != nil {
		id := ctx.Value("account").(*domain.Account).Id
		paymentLink = fmt.Sprintf("/pay/modify/%s?state=PAYED&id=%v", payId, id)
	}
	payment = &domain.Payment{
		PayState:    domain.PaymentState_WAITING,
		TotalPrice:  total,
		Expires:     DefaultProductFrozenExpires,
		CreateTime:  timestamppb.Now(),
		PayId:       payId,
		PaymentLink: paymentLink,
	}
	err = s.paymentRepo.Create(ctx, convertPayment(payment))
	if err != nil {
		return nil, err
	}
	settlementCache.SetDefault(payment.PayId, bill)
	klog.CtxDebugf(ctx, "创建支付订单，总额: %d", total)
	return
}

func (s *PaymentService) setupAutoThawedTrigger(ctx context.Context, payment *domain.Payment) (err error) {
	defer delete(lockMap, payment.PayId)
	l, ok := lockMap[payment.PayId]
	if !ok {
		l = &sync.Mutex{}
		lockMap[payment.PayId] = l
	}
	l.Lock()

	// 等待过期
	timer := time.NewTimer(time.Duration(payment.Expires) * time.Second)
	<-timer.C
	po, err := s.paymentRepo.GetByPayId(ctx, payment.PayId)
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return err
	}
	current := convertPo(po)
	if current.PayState == domain.PaymentState_WAITING {
		payment.PayState = domain.PaymentState_TIMEOUT
		err = s.paymentRepo.Update(ctx, convertPayment(payment))
		if err != nil {
			return err
		}
		klog.CtxInfof(ctx, "支付单%s当前状态为：WAITING，转变为：TIMEOUT", current.PayId)
		err = s.accomplishSettlement(ctx, domain.PaymentState_TIMEOUT, payment.PayId)
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
		}
	}
	timer.Stop()
	l.Unlock()
	return
}

func (s *PaymentService) accomplishSettlement(ctx context.Context, state domain.PaymentState, payId string) (err error) {
	settle, ok := settlementCache.Get(payId)
	if !ok {
		err = errors.New(fmt.Sprintf("支付单%s查询不到对应的支付结算单", payId))
		return
	}
	for _, item := range settle.(*domain.Settlement).Items {
		if state == domain.PaymentState_PAYED {
			err = warehouseClient.DecreaseStockpile(ctx, item.ProductId, item.Amount)
		} else {
			err = warehouseClient.ThawedStockpile(ctx, item.ProductId, item.Amount)
		}
		if err != nil {
			return
		}
	}
	return
}

func (s *PaymentService) rollbackSettlement(ctx context.Context, state domain.PaymentState, payId string) (err error) {
	settle, ok := settlementCache.Get(payId)
	if !ok {
		err = errors.New(fmt.Sprintf("支付单%s查询不到对应的支付结算单", payId))
		return
	}
	for _, item := range settle.(*domain.Settlement).Items {
		if state == domain.PaymentState_PAYED {
			err = warehouseClient.IncreaseStockpile(ctx, item.ProductId, item.Amount)
		} else {
			err = warehouseClient.FrozenStockpile(ctx, item.ProductId, item.Amount)
		}
		if err != nil {
			return
		}
	}
	return
}

func convertPo(po *PaymentPo) *domain.Payment {
	return &domain.Payment{
		PayId:       po.PayId,
		TotalPrice:  po.TotalPrice,
		Expires:     po.Expires,
		PaymentLink: po.PaymentLink,
		PayState:    domain.PaymentState(domain.PaymentState_value[po.PayState]),
		CreateTime:  timestamppb.New(po.CreateTime),
	}
}

func convertPayment(payment *domain.Payment) *PaymentPo {
	return &PaymentPo{
		PayId:       payment.PayId,
		CreateTime:  payment.CreateTime.AsTime(),
		Expires:     payment.Expires,
		TotalPrice:  payment.TotalPrice,
		PaymentLink: payment.PaymentLink,
		PayState:    domain.PaymentState_name[int32(payment.PayState)],
	}
}
