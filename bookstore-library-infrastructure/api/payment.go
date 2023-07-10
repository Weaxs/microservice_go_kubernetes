package api

import (
	"context"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
	"github.com/cloudwego/kitex/client/callopt"
)

// PaymentApi is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type PaymentApi interface {
	ExecuteSettlement(ctx context.Context, Req *domain.ExecuteSettlementRequest, callOptions ...callopt.Option) (r *domain.ExecuteSettlementResponse, err error)
	UpdatePaymentState(ctx context.Context, Req *domain.UpdatePaymentStateRequest, callOptions ...callopt.Option) (r *domain.EmptyResponse, err error)
	UpdatePaymentStateAlias(ctx context.Context, Req *domain.UpdatePaymentStateAlias, callOptions ...callopt.Option) (r *domain.EmptyResponse, err error)
}
