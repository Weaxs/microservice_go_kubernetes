package main

import (
	"context"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	account := &domain.Account{
		Username:  "zhangsan",
		Password:  "000000",
		Name:      "张三",
		Avatar:    "张三",
		Telephone: "13112341234",
		Email:     "zhangsan@gmail.com",
		Location:  "beijing",
	}
	_ = createAccount(context.Background(), account)

	err := createAccount(context.Background(), account)
	assert.Equal(t, "用户名称、邮箱、手机号码均不允许与现存用户重复", err.Error())
}

func TestUpdateAccount(t *testing.T) {
	account := &domain.Account{
		Username:  "zhangsan",
		Password:  "000000",
		Name:      "张三03",
		Avatar:    "张三",
		Telephone: "13112341234",
		Email:     "zhangsan@gmail.com",
		Location:  "beijing",
	}
	err := updateAccount(context.Background(), account)
	assert.Nil(t, err)
}
