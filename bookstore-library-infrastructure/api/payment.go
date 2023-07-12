package api

import (
	"context"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
)

type PaymentApi interface {
	ExecuteSettlement(ctx context.Context, req *domain.ExecuteSettlementRequest) (res *domain.ExecuteSettlementResponse, err error)
	UpdatePaymentState(ctx context.Context, req *domain.UpdatePaymentStateRequest) (res *domain.Empty, err error)
	UpdatePaymentStateAlias(ctx context.Context, req *domain.UpdatePaymentStateAlias) (res *domain.Empty, err error)
}
