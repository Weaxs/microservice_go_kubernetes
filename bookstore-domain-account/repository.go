package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db, _ = gorm.Open(mysql.Open("root:12345678@(127.0.0.1:3306)/bookstore"), &gorm.Config{})
)

type AccountDto struct {
	Id        int    `gorm:"primaryKey;autoIncrement:true" sql:"id"`
	Username  string `gorm:"index" sql:"username"`
	Password  string `gorm:"<-:create" sql:"password"`
	Name      string `sql:"name"`
	Avatar    string `sql:"avatar"`
	Telephone string `sql:"telephone"`
	Email     string `sql:"email"`
	Location  string `sql:"location"`
}

func (AccountDto) TableName() string {
	return "account"
}

func getByUsername(username string) (*AccountDto, error) {
	account := &AccountDto{}
	tx := db.First(account, "username = ?", username)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		return account, tx.Error
	}
}

func insert(account *AccountDto) error {
	result := db.Create(&account)
	return result.Error
}

func update(account *AccountDto) error {
	result := db.Model(account).Where("username = ?", account.Username).Updates(&account)
	return result.Error
}
