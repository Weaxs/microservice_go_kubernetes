package main

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db, _ = gorm.Open(mysql.Open("root:12345678@(127.0.0.1:3306)/bookstore"), &gorm.Config{})
)

type ProductRepo struct {
}
type StockpileRepo struct {
}
type SpecificationRepo struct {
}
type AdvertisementRepo struct {
}

type ProductPo struct {
	Id          int64   `gorm:"primaryKey;autoIncrement:true" sql:"id"`
	Title       string  `sql:"title"`
	Price       float64 `sql:"price"`
	Rate        float32 `sql:"rate"`
	Description string  `sql:"description"`
	Cover       string  `sql:"cover"`
	Detail      string  `sql:"detail"`
}

func (ProductPo) TableName() string {
	return "product"
}

type StockpilePo struct {
	Id        int64 `gorm:"primaryKey;autoIncrement:true" sql:"id"`
	Amount    int64 `gorm:"amount"`
	Frozen    int64 `gorm:"frozen"`
	ProductId int64 `gorm:"product_id"`
}

func (StockpilePo) TableName() string {
	return "stockpile"
}

type SpecificationPo struct {
	Id        int64  `gorm:"primaryKey;autoIncrement:true" sql:"id"`
	Item      string `sql:"item"`
	Value     string `sql:"value"`
	ProductId int64  `sql:"product_id"`
}

func (SpecificationPo) TableName() string {
	return "specification"
}

type AdvertisementPo struct {
	Id        int64  `gorm:"primaryKey;autoIncrement:true" sql:"id"`
	Image     string `sql:"image"`
	ProductId int64  `sql:"product_id"`
}

func (AdvertisementPo) TableName() string {
	return "advertisement"
}

func (*ProductRepo) FindAll(ctx context.Context) (products []*ProductPo, err error) {
	tx := db.Find(&products)
	if tx.Error != nil {
		err = tx.Error
	}
	return
}

func (*ProductRepo) FindById(ctx context.Context, productId int64) (product *ProductPo, err error) {
	tx := db.First(&product, "id = ?", productId)
	if tx.Error != nil {
		err = tx.Error
	}
	return
}

func (*ProductRepo) Create(ctx context.Context, product *ProductPo) (err error) {
	tx := db.Create(product)
	if tx.Error != nil {
		err = tx.Error
	}
	return
}

func (*ProductRepo) Update(ctx context.Context, product *ProductPo) (err error) {
	tx := db.Where("id = ?", product.Id).Updates(&product)
	if tx.Error != nil {
		err = tx.Error
	}
	return
}

func (*ProductRepo) DelById(ctx context.Context, productId int64) (err error) {
	tx := db.Where("id = ?", productId).Delete(&ProductPo{})
	if tx.Error != nil {
		err = tx.Error
	}
	return
}

func (*AdvertisementRepo) FindAll(ctx context.Context) (advertises []*AdvertisementPo, err error) {
	tx := db.Find(&advertises)
	if tx.Error != nil {
		err = tx.Error
	}
	return
}

func (*StockpileRepo) FindByProductId(ctx context.Context, productId int64) (s *StockpilePo, err error) {
	tx := db.First(&s, "product_id = ?", productId)
	if tx.Error != nil {
		err = tx.Error
	}
	return
}

func (*StockpileRepo) Update(ctx context.Context, stockpile *StockpilePo) (err error) {
	tx := db.Model(stockpile).Where("product_id = ?", stockpile.ProductId).Updates(stockpile)
	if tx.Error != nil {
		err = tx.Error
	}
	return
}

func (*SpecificationRepo) Create(ctx context.Context, specifications []*SpecificationPo) (err error) {
	tx := db.Create(specifications)
	if tx.Error != nil {
		err = tx.Error
	}
	return
}

func (*SpecificationRepo) FindByProductId(ctx context.Context, productId int64) (specifications []*SpecificationPo, err error) {
	tx := db.Find(&specifications, "product_id = ?", productId)
	if tx.Error != nil {
		err = tx.Error
	}
	return
}

func (*SpecificationRepo) DelByProductId(ctx context.Context, productId int64) (err error) {
	tx := db.Where("product_id = ?", productId).Delete(&SpecificationPo{})
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}
