package main

import (
	"context"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/client"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
	kitex "github.com/cloudwego/kitex/client"
)

type WarehouseClient struct {
	client client.WarehouseClient
}

func NewWarehouseClient() (*WarehouseClient, error) {
	c, err := client.NewWarehouseClient(
		kitex.WithHostPorts("[::1]:8811"), kitex.WithMuxConnection(1))
	if err != nil {
		return nil, err
	}
	return &WarehouseClient{
		client: c,
	}, nil
}

func (c *WarehouseClient) GetProducts(ctx context.Context) (products []*domain.Product, err error) {
	resp, err := c.client.GetAllProducts(ctx, &domain.EmptyResponse{})
	if err != nil {
		return nil, err
	}
	return resp.GetProducts(), nil
}

func (c *WarehouseClient) FrozenStockpile(ctx context.Context, productId, amount int64) (err error) {
	req := &domain.SetDeliveredStatusRequest{ProductId: productId, Amount: amount, Status: domain.DeliveredStatus_FROZEN}
	_, err = c.client.SetDeliveredStatus(ctx, req)
	if err != nil {
		return
	}
	return
}

func (c *WarehouseClient) IncreaseStockpile(ctx context.Context, productId, amount int64) (err error) {
	req := &domain.SetDeliveredStatusRequest{ProductId: productId, Amount: amount, Status: domain.DeliveredStatus_INCREASE}
	_, err = c.client.SetDeliveredStatus(ctx, req)
	if err != nil {
		return
	}
	return
}

func (c *WarehouseClient) DecreaseStockpile(ctx context.Context, productId, amount int64) (err error) {
	req := &domain.SetDeliveredStatusRequest{ProductId: productId, Amount: amount, Status: domain.DeliveredStatus_DECREASE}
	_, err = c.client.SetDeliveredStatus(ctx, req)
	if err != nil {
		return
	}
	return
}

func (c *WarehouseClient) ThawedStockpile(ctx context.Context, productId, amount int64) (err error) {
	req := &domain.SetDeliveredStatusRequest{ProductId: productId, Amount: amount, Status: domain.DeliveredStatus_THAWED}
	_, err = c.client.SetDeliveredStatus(ctx, req)
	if err != nil {
		return
	}
	return
}
