package main

import (
	"context"
	"fmt"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
	"testing"
)

func TestExecuteSettlement(t *testing.T) {
	var items []*domain.Item
	items = append(items, &domain.Item{ProductId: 1, Amount: 1})
	items = append(items, &domain.Item{ProductId: 2, Amount: 1})
	settle := &domain.Settlement{
		Items: items,
		Purchase: &domain.Purchase{
			Location: "xx rd. zhuhai. guangdong. china", Name: "icyfenix", Telephone: "18888888888"},
	}
	payment, err := paymentService.ExecuteBySettlement(context.Background(), settle)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(payment)
}
