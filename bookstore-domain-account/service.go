package main

import (
	"context"
	"errors"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
	"github.com/asaskevich/govalidator"
	"github.com/cloudwego/kitex/pkg/klog"
)

func getAccount(ctx context.Context, username string) (account *domain.Account, err error) {
	if govalidator.IsNull(username) {
		klog.CtxErrorf(ctx, "get account error: Username is blank")
		return nil, errors.New("username is blank")
	}

	dto, err := getByUsername(username)
	if err != nil {
		return nil, err
	}

	klog.CtxDebugf(ctx, "get from db")
	return &domain.Account{
		Username:  dto.Username,
		Password:  dto.Password,
		Name:      dto.Name,
		Avatar:    dto.Avatar,
		Telephone: dto.Telephone,
		Email:     dto.Email,
		Location:  dto.Location,
	}, nil
}

func createAccount(ctx context.Context, account *domain.Account) error {
	if _, err := govalidator.ValidateStruct(account); err != nil {
		return err
	}

	dto := &AccountDto{
		Username:  account.Username,
		Password:  account.Password,
		Name:      account.Name,
		Avatar:    account.Avatar,
		Telephone: account.Telephone,
		Email:     account.Email,
		Location:  account.Location,
	}
	klog.CtxDebugf(ctx, "insert into db")
	err := insert(dto)
	if err != nil {
		return err
	}

	return nil
}

func updateAccount(ctx context.Context, account *domain.Account) error {
	if _, err := govalidator.ValidateStruct(account); err != nil {
		return err
	}

	dto := &AccountDto{
		Username:  account.Username,
		Password:  account.Password,
		Name:      account.Name,
		Avatar:    account.Avatar,
		Telephone: account.Telephone,
		Email:     account.Email,
		Location:  account.Location,
	}
	klog.CtxDebugf(ctx, "update from db")
	err := update(dto)
	if err != nil {
		return err
	}

	return nil
}
