package api

import (
	"context"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
)

type AccountApi interface {
	GetAccount(ctx context.Context, username *domain.GetAccountRequest) (r *domain.GetAccountResponse, err error)
	CreateAccount(ctx context.Context, account *domain.ChangeAccountRequest) (r *domain.ChangeAccountResponse, err error)
	UpdateAccount(ctx context.Context, account *domain.ChangeAccountRequest) (r *domain.ChangeAccountResponse, err error)
}
