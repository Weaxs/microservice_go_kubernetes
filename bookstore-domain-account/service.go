package main

import (
	"context"
	"errors"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
	"github.com/asaskevich/govalidator"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
)

func getAccount(ctx context.Context, username string) (account *domain.Account, err error) {
	if govalidator.IsNull(username) {
		klog.CtxErrorf(ctx, "get account error: Username is blank")
		return nil, errors.New("username is blank")
	}

	po, err := getByUsername(username)
	if err != nil {
		return nil, err
	}
	if po == nil {
		return nil, errors.New("用户不存在")
	}

	klog.CtxDebugf(ctx, "get from db")
	return &domain.Account{
		Id:        po.Id,
		Username:  po.Username,
		Password:  po.Password,
		Name:      po.Name,
		Avatar:    po.Avatar,
		Telephone: po.Telephone,
		Email:     po.Email,
		Location:  po.Location,
	}, nil
}

func createAccount(ctx context.Context, account *domain.Account) (err error) {
	if _, err := govalidator.ValidateStruct(account); err != nil {
		return err
	}

	exist, _ := findByUsernameOrEmailOrTelephone(account.Username, account.Email, account.Telephone)
	if exist != nil {
		klog.CtxDebugf(ctx, "用户名称、邮箱、手机号码均不允许与现存用户重复")
		return errors.New("用户名称、邮箱、手机号码均不允许与现存用户重复")
	}
	// Bcrypt加密
	password, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	account.Password = string(password)

	po := &AccountPo{
		Username:  account.Username,
		Password:  account.Password,
		Name:      account.Name,
		Avatar:    account.Avatar,
		Telephone: account.Telephone,
		Email:     account.Email,
		Location:  account.Location,
	}
	klog.CtxDebugf(ctx, "insert into db")
	err = insert(po)
	if err != nil {
		return err
	}

	return nil
}

func updateAccount(ctx context.Context, account *domain.Account) (err error) {
	if _, err := govalidator.ValidateStruct(account); err != nil {
		return err
	}
	origin, err := getByUsername(account.Username)

	exist, _ := findByUsernameOrEmailOrTelephone(account.Username, account.Email, account.Telephone)
	if exist != nil && exist.Id != origin.Id {
		klog.CtxDebugf(ctx, "用户名称、邮箱、手机号码与现存用户产生重复")
		return errors.New("用户名称、邮箱、手机号码与现存用户产生重复")
	}

	po := &AccountPo{
		Username:  account.Username,
		Password:  account.Password,
		Name:      account.Name,
		Avatar:    account.Avatar,
		Telephone: account.Telephone,
		Email:     account.Email,
		Location:  account.Location,
	}
	klog.CtxDebugf(ctx, "update from db")
	err = update(po)
	if err != nil {
		return err
	}

	return nil
}
