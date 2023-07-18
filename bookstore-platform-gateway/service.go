package main

import (
	"context"
	"encoding/json"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/client"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

type WarehouseService struct {
	client client.WarehouseClient
}

type AccountService struct {
	client client.AccountClient
}

type PaymentService struct {
	client client.PaymentClient
}

func NewAccountService() *AccountService {
	c, _ := NewAccountClient()
	return &AccountService{client: c}
}

func NewWarehouseService() *WarehouseService {
	c, _ := NewWarehouseClient()
	return &WarehouseService{client: c}
}

func NewPaymentService() *PaymentService {
	c, _ := NewPaymentClient()
	return &PaymentService{client: c}
}

func (s *AccountService) findAccountByUsername(c context.Context, ctx *app.RequestContext) {
	req := &domain.GetAccountRequest{Username: ctx.Param("username")}
	resp, err := s.client.GetAccount(c, req)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	ctx.JSON(consts.StatusOK, resp.Account)
}

func (s *AccountService) createAccount(c context.Context, ctx *app.RequestContext) {
	body, err := ctx.Body()
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	var account *domain.Account
	err = json.Unmarshal(body, &account)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}

	req := &domain.ChangeAccountRequest{Account: account}
	_, err = s.client.CreateAccount(c, req)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	ctx.Status(consts.StatusOK)
}

func (s *AccountService) updateAccount(c context.Context, ctx *app.RequestContext) {
	body, err := ctx.Body()
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	var account *domain.Account
	err = json.Unmarshal(body, &account)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}

	req := &domain.ChangeAccountRequest{Account: account}
	_, err = s.client.UpdateAccount(c, req)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	ctx.Status(consts.StatusOK)
}

func (s *PaymentService) executeSettlement(c context.Context, ctx *app.RequestContext) {
	body, err := ctx.Body()
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	var settlement *domain.Settlement
	err = json.Unmarshal(body, &settlement)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}

	req := &domain.ExecuteSettlementRequest{Settlement: settlement}
	resp, err := s.client.ExecuteSettlement(c, req)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	ctx.JSON(consts.StatusOK, resp.Payment)
}

func (s *PaymentService) modifyPaymentState(c context.Context, ctx *app.RequestContext) {
	payId := ctx.Param("id")
	state := ctx.Param("state")
	req := &domain.UpdatePaymentStateAlias{PayId: payId, State: domain.PaymentState(domain.PaymentState_value[state])}
	_, err := s.client.UpdatePaymentStateAlias(c, req)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	ctx.Status(consts.StatusOK)
}

func (s *WarehouseService) getAllProducts(c context.Context, ctx *app.RequestContext) {
	resp, err := s.client.GetAllProducts(c, &domain.Empty{})
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	ctx.JSON(consts.StatusOK, resp.Products)
}

func (s *WarehouseService) getProductById(c context.Context, ctx *app.RequestContext) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	resp, err := s.client.GetProduct(c, &domain.GetProductRequest{Id: int64(id)})
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	ctx.JSON(consts.StatusOK, resp.Product)
}

func (s *WarehouseService) createProduct(c context.Context, ctx *app.RequestContext) {
	body, err := ctx.Body()
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	var product *domain.Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}

	req := &domain.ChangeProductRequest{Product: product}
	_, err = s.client.CreateProduct(c, req)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	ctx.Status(consts.StatusOK)
}

func (s *WarehouseService) updateProduct(c context.Context, ctx *app.RequestContext) {
	body, err := ctx.Body()
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	var product *domain.Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}

	req := &domain.ChangeProductRequest{Product: product}
	_, err = s.client.UpdateProduct(c, req)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	ctx.Status(consts.StatusOK)
}

func (s *WarehouseService) deleteById(c context.Context, ctx *app.RequestContext) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}

	req := &domain.RemoveProductRequest{Id: int64(id)}
	_, err = s.client.RemoveProduct(c, req)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	ctx.Status(consts.StatusOK)
}

func (s *WarehouseService) getStockpileByProductId(c context.Context, ctx *app.RequestContext) {
	idStr := ctx.Param("productId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}

	req := &domain.QueryStockpileRequest{ProductId: int64(id)}
	resp, err := s.client.QueryStockpile(c, req)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	ctx.JSON(consts.StatusOK, resp.Stockpile)
}

func (s *WarehouseService) updateStockpile(c context.Context, ctx *app.RequestContext) {
	idStr := ctx.Param("productId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	amountStr := ctx.Param("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}

	req := &domain.UpdateStockpileRequest{ProductId: int64(id), Amount: int64(amount)}
	_, err = s.client.UpdateStockpile(c, req)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	ctx.Status(consts.StatusOK)
}
func (s *WarehouseService) getAllAdvertisements(c context.Context, ctx *app.RequestContext) {
	resp, err := s.client.GetAllAdvertisements(c, &domain.Empty{})
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}
	ctx.JSON(consts.StatusOK, resp.Advertisements)
}
