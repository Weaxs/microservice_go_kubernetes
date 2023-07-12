package main

import (
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/client"
	kitex "github.com/cloudwego/kitex/client"
)

var (
	WarehouseClient, _ = NewWarehouseClient()
	AccountClient, _   = NewAccountClient()
	PaymentClient, _   = NewPaymentClient()
)

func NewWarehouseClient() (*client.WarehouseClient, error) {
	c, err := client.NewWarehouseClient(
		kitex.WithHostPorts("[::1]:8811"), kitex.WithMuxConnection(1))
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func NewAccountClient() (*client.AccountClient, error) {
	c, err := client.NewAccountClient(
		kitex.WithHostPorts("[::1]:8811"), kitex.WithMuxConnection(1))
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func NewPaymentClient() (*client.PaymentClient, error) {
	c, err := client.NewPaymentClient(
		kitex.WithHostPorts("[::1]:8811"), kitex.WithMuxConnection(1))
	if err != nil {
		return nil, err
	}
	return &c, nil
}
