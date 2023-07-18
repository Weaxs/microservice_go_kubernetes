package main

import (
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/client"
	kitex "github.com/cloudwego/kitex/client"
)

func NewWarehouseClient() (client.WarehouseClient, error) {
	hostports := Config.GetStringSlice(WarehouseClientHostPost)
	c, err := client.NewWarehouseClient(
		kitex.WithHostPorts(hostports...), kitex.WithMuxConnection(Config.GetInt(WarehouseClientConnNum)))
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewAccountClient() (client.AccountClient, error) {
	hostports := Config.GetStringSlice(AccountClientHostPost)
	c, err := client.NewAccountClient(
		kitex.WithHostPorts(hostports...), kitex.WithMuxConnection(Config.GetInt(AccountClientConnNum)))
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewPaymentClient() (client.PaymentClient, error) {
	hostports := Config.GetStringSlice(PaymentClientHostPost)
	c, err := client.NewPaymentClient(
		kitex.WithHostPorts(hostports...), kitex.WithMuxConnection(Config.GetInt(PaymentClientConnNum)))
	if err != nil {
		return nil, err
	}
	return c, nil
}
