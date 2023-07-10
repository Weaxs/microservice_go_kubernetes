package main

import (
	"context"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
)

// handler implements the last service interface defined in the IDL.
type handler struct{}

// ExecuteSettlement implements the handler interface.
func (s *handler) ExecuteSettlement(ctx context.Context, req *domain.ExecuteSettlementRequest) (resp *domain.ExecuteSettlementResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdatePaymentState implements the handler interface.
func (s *handler) UpdatePaymentState(ctx context.Context, req *domain.UpdatePaymentStateRequest) (resp *domain.EmptyResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdatePaymentStateAlias implements the handler interface.
func (s *handler) UpdatePaymentStateAlias(ctx context.Context, req *domain.UpdatePaymentStateAlias) (resp *domain.EmptyResponse, err error) {
	// TODO: Your code here...

	return
}
