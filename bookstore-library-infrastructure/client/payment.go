package client

import (
	"context"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/api"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// NewPaymentClient creates a client for the service defined in IDL.
func NewPaymentClient(opts ...client.Option) (api.PaymentApi, error) {
	var options []client.Option
	options = append(options, client.WithDestService("PaymentApi"))

	options = append(options, opts...)

	kc, err := client.NewClient(api.PaymentApiServiceInfo, options...)
	if err != nil {
		return nil, err
	}
	return &kPaymentApiClient{
		kClient: newServiceClient(kc),
	}, nil
}

type kPaymentApiClient struct {
	*kClient
}

func (p *kPaymentApiClient) ExecuteSettlement(ctx context.Context, Req *domain.ExecuteSettlementRequest, callOptions ...callopt.Option) (r *domain.ExecuteSettlementResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ExecuteSettlement(ctx, Req)
}

func (p *kPaymentApiClient) UpdatePaymentState(ctx context.Context, Req *domain.UpdatePaymentStateRequest, callOptions ...callopt.Option) (r *domain.EmptyResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatePaymentState(ctx, Req)
}

func (p *kPaymentApiClient) UpdatePaymentStateAlias(ctx context.Context, Req *domain.UpdatePaymentStateAlias, callOptions ...callopt.Option) (r *domain.EmptyResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatePaymentStateAlias(ctx, Req)
}

func (p *kClient) ExecuteSettlement(ctx context.Context, Req *domain.ExecuteSettlementRequest) (r *domain.ExecuteSettlementResponse, err error) {
	var _args domain.ExecuteSettlementArgs
	_args.Req = Req
	var _result domain.ExecuteSettlementResult
	if err = p.c.Call(ctx, "executeSettlement", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdatePaymentState(ctx context.Context, Req *domain.UpdatePaymentStateRequest) (r *domain.EmptyResponse, err error) {
	var _args domain.UpdatePaymentStateArgs
	_args.Req = Req
	var _result domain.UpdatePaymentStateResult
	if err = p.c.Call(ctx, "updatePaymentState", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdatePaymentStateAlias(ctx context.Context, Req *domain.UpdatePaymentStateAlias) (r *domain.EmptyResponse, err error) {
	var _args domain.UpdatePaymentStateAliasArgs
	_args.Req = Req
	var _result domain.UpdatePaymentStateAliasResult
	if err = p.c.Call(ctx, "updatePaymentStateAlias", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
