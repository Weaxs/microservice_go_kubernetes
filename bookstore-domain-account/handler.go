package main

import (
	"context"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
	"github.com/cloudwego/kitex/pkg/klog"
)

type handler struct{}

func (h *handler) GetAccount(ctx context.Context, username *domain.GetAccountRequest) (r *domain.GetAccountResponse, err error) {
	klog.CtxDebugf(ctx, "get user by Username: %s", username)

	account, err := getAccount(ctx, username.Username)
	if err != nil {
		return nil, err
	}
	return &domain.GetAccountResponse{Account: account}, nil
}

func (h *handler) CreateAccount(ctx context.Context, account *domain.ChangeAccountRequest) (r *domain.ChangeAccountResponse, err error) {
	klog.CtxDebugf(ctx, "create user")
	err = createAccount(ctx, account.Account)
	if err != nil {
		return nil, err
	}
	return &domain.ChangeAccountResponse{}, nil
}

func (h *handler) UpdateAccount(ctx context.Context, account *domain.ChangeAccountRequest) (r *domain.ChangeAccountResponse, err error) {
	klog.CtxDebugf(ctx, "update user")
	err = updateAccount(ctx, account.Account)
	if err != nil {
		return nil, err
	}
	return &domain.ChangeAccountResponse{}, nil
}
