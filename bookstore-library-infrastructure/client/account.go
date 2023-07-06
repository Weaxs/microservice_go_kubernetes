package client

import (
	"context"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/api"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// AccountClient is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type AccountClient interface {
	GetAccount(ctx context.Context, usernmae *domain.GetAccountRequest, callOptions ...callopt.Option) (r *domain.GetAccountResponse, err error)
	CreateAccount(ctx context.Context, account *domain.ChangeAccountRequest, callOptions ...callopt.Option) (r *domain.ChangeAccountResponse, err error)
	UpdateAccount(ctx context.Context, account *domain.ChangeAccountRequest, callOptions ...callopt.Option) (r *domain.ChangeAccountResponse, err error)
}

// NewAccountClient creates a client for the service defined in IDL.
func NewAccountClient(opts ...client.Option) (AccountClient, error) {
	var options []client.Option
	options = append(options, client.WithDestService("AccountApi"))

	options = append(options, opts...)

	kc, err := client.NewClient(api.AccountApiServiceInfo, options...)
	if err != nil {
		return nil, err
	}
	return &kAccountApiClient{
		kClient: newServiceClient(kc),
	}, nil
}

type kAccountApiClient struct {
	*kClient
}

func (p *kAccountApiClient) GetAccount(ctx context.Context, usernmae *domain.GetAccountRequest, callOptions ...callopt.Option) (r *domain.GetAccountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAccount(ctx, usernmae)
}

func (p *kAccountApiClient) CreateAccount(ctx context.Context, account *domain.ChangeAccountRequest, callOptions ...callopt.Option) (r *domain.ChangeAccountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateAccount(ctx, account)
}

func (p *kAccountApiClient) UpdateAccount(ctx context.Context, account *domain.ChangeAccountRequest, callOptions ...callopt.Option) (r *domain.ChangeAccountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateAccount(ctx, account)
}

func (p *kClient) GetAccount(ctx context.Context, usernmae *domain.GetAccountRequest) (r *domain.GetAccountResponse, err error) {
	var _args domain.GetAccountArgs
	_args.Username = usernmae
	var _result domain.GetAccountResult
	if err = p.c.Call(ctx, "getAccount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateAccount(ctx context.Context, account *domain.ChangeAccountRequest) (r *domain.ChangeAccountResponse, err error) {
	var _args domain.ChangeAccountArgs
	_args.Account = account
	var _result domain.ChangeAccountResult
	if err = p.c.Call(ctx, "createAccount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateAccount(ctx context.Context, account *domain.ChangeAccountRequest) (r *domain.ChangeAccountResponse, err error) {
	var _args domain.ChangeAccountArgs
	_args.Account = account
	var _result domain.ChangeAccountResult
	if err = p.c.Call(ctx, "updateAccount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
