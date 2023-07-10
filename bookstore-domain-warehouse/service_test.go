package main

import (
	"context"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdvertiseService_GetAll(t *testing.T) {
	all, err := advertiseService.GetAll(context.Background())
	if err != nil {
		assert.Fail(t, err.Error())
	} else {
		assert.Equal(t, 3, len(all))
	}
}

func TestProductService_GetAll(t *testing.T) {
	all, err := productService.GetAll(context.Background())
	if err != nil {
		assert.Fail(t, err.Error())
		return
	} else {
		assert.Equal(t, 8, len(all))
	}
}

func TestProductService_GetById(t *testing.T) {
	product, err := productService.GetById(context.Background(), int64(1))
	if err != nil {
		assert.Fail(t, err.Error())
		return
	} else {
		assert.Equal(t, "深入理解Java虚拟机（第3版）", product.Title)
	}
}

func TestProductService_Create(t *testing.T) {
	var specifications []*domain.Specification
	specifications = append(specifications, &domain.Specification{Item: "作者", Value: "厄休拉·勒古恩"})
	specifications = append(specifications, &domain.Specification{Item: "译者", Value: "夏笳"})
	specifications = append(specifications, &domain.Specification{Item: "原作名", Value: "Words Are My Matter"})
	specifications = append(specifications, &domain.Specification{Item: "装帧", Value: "精装"})
	specifications = append(specifications, &domain.Specification{Item: "页数", Value: "536"})
	specifications = append(specifications, &domain.Specification{Item: "出版年", Value: "2023-5"})
	specifications = append(specifications, &domain.Specification{Item: "ISBN", Value: "9787555914105"})
	product := &domain.Product{
		Title: "我以文字为业",
		Price: float64(72),
		Rate:  float32(10),
		Description: "<p>阅读 写作 批评 行动</p>" +
			"<p>六十八篇散文 一部“思想自传”</p>" +
			"<p>“艰难的时代要来了，在那样的时代里，我们将会需要另一些作家的声音，他们能够看到与我们当下不同的生活方式，能够穿过我们饱受恐惧之苦的社会，穿过其对技术的痴迷，去看到其他生存道路，甚至能够想象希望的真正土壤。</p>" +
			"<p>我们将会需要能够记住自由的作家——诗人，富有远见的人——能够把握一种更大现实的现实主义者。”</p>",
		Cover:          "",
		Detail:         "",
		Specifications: specifications,
	}
	err := productService.Create(context.Background(), product)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}
}

func TestProductService_RemoveById(t *testing.T) {
	err := productService.RemoveById(context.Background(), int64(9))
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}
}

func TestStockpileService_QueryByProduct(t *testing.T) {
	stockpile, err := stockpileService.QueryByProduct(context.Background(), int64(1))
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}
	assert.Equal(t, int64(1), stockpile.Product.Id)
}

func TestStockpileService_Update(t *testing.T) {
	err := stockpileService.Update(context.Background(), 11, 50)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}
}

func TestStockpileService_SetDeliveredStatus(t *testing.T) {
	err := stockpileService.SetDeliveredStatus(context.Background(), 11, 10, domain.DeliveredStatus_DECREASE)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}
}
