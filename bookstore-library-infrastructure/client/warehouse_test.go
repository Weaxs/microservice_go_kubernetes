package client

import (
	"context"
	"fmt"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
	"github.com/cloudwego/kitex/client"
	"testing"
)

func TestWarehouseClient(t *testing.T) {
	client, err := NewWarehouseClient(
		client.WithHostPorts("[::1]:8811"),
		client.WithMuxConnection(1))
	if err != nil {
		panic(any(err))
	}
	products, err := client.GetAllProducts(context.Background(), &domain.Empty{})
	if err != nil {
		panic(any(err))
	}
	fmt.Print(products)
}
