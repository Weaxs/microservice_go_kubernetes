package main

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	dsn = Config.GetString(DbUserKey) + ":" + Config.GetString(DbPasswordKey) + "@(" +
		Config.GetString(DbHostKey) + ":" + Config.GetString(DbPortKey) + ")/" +
		Config.GetString(DbDatabaseKey) + "?parseTime=true"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
)

type WalletRepo struct {
}
type PaymentRepo struct {
}

type WalletPo struct {
	Id        int64   `gorm:"primaryKey;autoIncrement:true" sql:"id"`
	Money     float64 `sql:"money"`
	AccountId int64   `sql:"account_id"`
}

func (WalletPo) TableName() string {
	return "wallet"
}

type PaymentPo struct {
	Id          int64     `gorm:"primaryKey;autoIncrement:true" sql:"id"`
	PayId       string    `sql:"pay_id"`
	CreateTime  time.Time `sql:"create_time"`
	TotalPrice  float64   `sql:"total_price"`
	Expires     int64     `sql:"expires"`
	PaymentLink string    `sql:"payment_link"`
	PayState    string    `sql:"pay_state"`
}

func (PaymentPo) TableName() string {
	return "payment"
}

func (r *WalletRepo) FindByAccountId(ctx context.Context, accountId int64) (wallet *WalletPo, err error) {
	tx := db.First(&wallet, "account_id = ?", accountId)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}

func (r *WalletRepo) Save(ctx context.Context, wallet *WalletPo) (err error) {
	tx := db.Save(&wallet)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}

func (r *PaymentRepo) GetByPayId(ctx context.Context, payId string) (payment *PaymentPo, err error) {
	tx := db.First(&payment, "pay_id = ?", payId)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}

func (r *PaymentRepo) Create(ctx context.Context, payment *PaymentPo) (err error) {
	tx := db.Create(&payment)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}

func (r *PaymentRepo) Update(ctx context.Context, payment *PaymentPo) (err error) {
	tx := db.Where("pay_id = ?", payment.PayId).Updates(&payment)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}
