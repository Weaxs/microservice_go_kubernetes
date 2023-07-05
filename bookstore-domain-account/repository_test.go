package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepo(t *testing.T) {
	dto := &AccountDto{
		Username:  "zhangsan",
		Password:  "000000",
		Name:      "张三",
		Avatar:    "张三",
		Telephone: "13112341234",
		Email:     "zhangsan@gmail.com",
		Location:  "beijing",
	}
	testInsert(t, dto)

	ret := testGetByUsername(t, "zhangsan")
	if ret == nil {
		t.Error(errors.New("can not found zhangsan"))
		t.FailNow()
	}

	dto.Name = "张三03"
	dto.Password = "111111"
	testUpdate(t, dto)
	ret = testGetByUsername(t, "zhangsan")
	if ret == nil {
		t.Error(errors.New("can not found zhangsan"))
		t.FailNow()
	}

	assert.Equal(t, ret.Password, "000000")
	assert.Equal(t, ret.Name, "张三03")
}

func testInsert(t *testing.T, dto *AccountDto) {
	err := insert(dto)
	if err != nil {
		t.Error(err)
	}
}

func testUpdate(t *testing.T, dto *AccountDto) {
	err := update(dto)
	if err != nil {
		t.Error(err)
	}
}

func testGetByUsername(t *testing.T, username string) *AccountDto {
	ret, err := getByUsername(username)
	if err != nil {
		t.Error(err)
	}
	return ret
}
