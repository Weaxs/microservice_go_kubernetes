package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dsn = Config.GetString(DbUserKey) + ":" + Config.GetString(DbPasswordKey) + "@(" +
		Config.GetString(DbHostKey) + ":" + Config.GetString(DbPortKey) + ")/" +
		Config.GetString(DbDatabaseKey) + "?parseTime=true"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
)

type AccountPo struct {
	Id        int64  `gorm:"primaryKey;autoIncrement:true" sql:"id"`
	Username  string `gorm:"index" sql:"username"`
	Password  string `gorm:"<-:create" sql:"password"`
	Name      string `sql:"name"`
	Avatar    string `sql:"avatar"`
	Telephone string `sql:"telephone"`
	Email     string `sql:"email"`
	Location  string `sql:"location"`
}

func (AccountPo) TableName() string {
	return "account"
}

func getByUsername(username string) (*AccountPo, error) {
	account := &AccountPo{}
	tx := db.First(account, "username = ?", username)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		return account, tx.Error
	}
}

func insert(account *AccountPo) error {
	result := db.Create(&account)
	return result.Error
}

func update(account *AccountPo) error {
	result := db.Model(account).Where("username = ?", account.Username).Updates(&account)
	return result.Error
}

func findByUsernameOrEmailOrTelephone(username, email, telephone string) (*AccountPo, error) {
	account := &AccountPo{}
	tx := db.Where("username = ? ", username).Or("email = ?", email).
		Or("telephone = ?", telephone).First(account)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return account, nil
}
