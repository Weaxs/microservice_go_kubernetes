package client

import (
	"context"
	"fmt"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
	"github.com/cloudwego/kitex/client"
	"testing"
)

func TestPaymentClient(t *testing.T) {
	c, err := NewPaymentClient(
		client.WithHostPorts("[::1]:8812"),
		client.WithMuxConnection(1))
	if err != nil {
		panic(any(err))
	}
	payment, err := c.ExecuteSettlement(context.Background(), &domain.ExecuteSettlementRequest{})
	if err != nil {
		panic(any(err))
	}
	fmt.Print(payment)
}
