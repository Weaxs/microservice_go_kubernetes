package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepo(t *testing.T) {
	po := &AccountPo{
		Username:  "zhangsan",
		Password:  "000000",
		Name:      "张三",
		Avatar:    "张三",
		Telephone: "13112341234",
		Email:     "zhangsan@gmail.com",
		Location:  "beijing",
	}
	testInsert(t, po)

	ret := testGetByUsername(t, "zhangsan")
	if ret == nil {
		t.Error(errors.New("can not found zhangsan"))
		t.FailNow()
	}

	po.Name = "张三03"
	po.Password = "111111"
	testUpdate(t, po)
	ret = testGetByUsername(t, "zhangsan")
	if ret == nil {
		t.Error(errors.New("can not found zhangsan"))
		t.FailNow()
	}

	assert.Equal(t, ret.Password, "000000")
	assert.Equal(t, ret.Name, "张三03")
}

func testInsert(t *testing.T, dto *AccountPo) {
	err := insert(dto)
	if err != nil {
		t.Error(err)
	}
}

func testUpdate(t *testing.T, po *AccountPo) {
	err := update(po)
	if err != nil {
		t.Error(err)
	}
}

func testGetByUsername(t *testing.T, username string) *AccountPo {
	ret, err := getByUsername(username)
	if err != nil {
		t.Error(err)
	}
	return ret
}
