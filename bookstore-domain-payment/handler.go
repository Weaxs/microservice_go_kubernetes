package main

import (
	"context"
	"errors"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
)

// handler implements the last service interface defined in the IDL.
type handler struct{}

// ExecuteSettlement implements the handler interface.
func (s *handler) ExecuteSettlement(ctx context.Context, req *domain.ExecuteSettlementRequest) (resp *domain.ExecuteSettlementResponse, err error) {
	payment, err := paymentService.ExecuteBySettlement(ctx, req.Settlement)
	if err != nil {
		return nil, err
	}
	return &domain.ExecuteSettlementResponse{Payment: payment}, nil
}

// UpdatePaymentState implements the handler interface.
func (s *handler) UpdatePaymentState(ctx context.Context, req *domain.UpdatePaymentStateRequest) (resp *domain.EmptyResponse, err error) {
	account := ctx.Value("account")
	if account == nil {
		return nil, errors.New("there is no account in context")
	}
	return s.UpdatePaymentStateAlias(ctx, &domain.UpdatePaymentStateAlias{
		State: req.State, PayId: req.PayId, AccountId: account.(*domain.Account).Id})
}

// UpdatePaymentStateAlias implements the handler interface.
func (s *handler) UpdatePaymentStateAlias(ctx context.Context, req *domain.UpdatePaymentStateAlias) (resp *domain.EmptyResponse, err error) {
	state := req.State
	if state == domain.PaymentState_PAYED {
		err = paymentService.AccomplishPayment(ctx, req.PayId, req.AccountId)
	} else {
		err = paymentService.CancelPayment(ctx, req.PayId)
	}
	return
}
