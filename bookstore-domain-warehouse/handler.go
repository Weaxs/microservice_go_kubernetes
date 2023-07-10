package main

import (
	"context"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
)

// handler implements the last service interface defined in the IDL.
type handler struct{}

// GetAllAdvertisements implements the handler interface.
func (s *handler) GetAllAdvertisements(ctx context.Context, req *domain.EmptyResponse) (resp *domain.GetAllAdvertisementsResponse, err error) {
	advertises, err := advertiseService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return &domain.GetAllAdvertisementsResponse{Advertisements: advertises}, nil
}

// GetAllProducts implements the handler interface.
func (s *handler) GetAllProducts(ctx context.Context, req *domain.EmptyResponse) (resp *domain.GetAllProductResponse, err error) {
	products, err := productService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return &domain.GetAllProductResponse{Products: products}, nil
}

// GetProduct implements the handler interface.
func (s *handler) GetProduct(ctx context.Context, req *domain.GetProductRequest) (resp *domain.GetProductResponse, err error) {
	product, err := productService.GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &domain.GetProductResponse{Product: product}, nil
}

// CreateProduct implements the handler interface.
func (s *handler) CreateProduct(ctx context.Context, req *domain.ChangeProductRequest) (resp *domain.EmptyResponse, err error) {
	err = productService.Create(ctx, req.Product)
	if err != nil {
		return nil, err
	}
	return &domain.EmptyResponse{}, nil
}

// UpdateProduct implements the handler interface.
func (s *handler) UpdateProduct(ctx context.Context, req *domain.ChangeProductRequest) (resp *domain.EmptyResponse, err error) {
	err = productService.Update(ctx, req.Product)
	if err != nil {
		return nil, err
	}
	return &domain.EmptyResponse{}, nil
}

// RemoveProduct implements the handler interface.
func (s *handler) RemoveProduct(ctx context.Context, req *domain.RemoveProductRequest) (resp *domain.EmptyResponse, err error) {
	err = productService.RemoveById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &domain.EmptyResponse{}, nil
}

// UpdateStockpile implements the handler interface.
func (s *handler) UpdateStockpile(ctx context.Context, req *domain.UpdateStockpileRequest) (resp *domain.EmptyResponse, err error) {
	err = stockpileService.Update(ctx, req.ProductId, req.Amount)
	if err != nil {
		return nil, err
	}
	return &domain.EmptyResponse{}, nil
}

// QueryStockpile implements the handler interface.
func (s *handler) QueryStockpile(ctx context.Context, req *domain.QueryStockpileRequest) (resp *domain.QueryStockpileResponse, err error) {
	stockpile, err := stockpileService.QueryByProduct(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}
	return &domain.QueryStockpileResponse{Stockpile: stockpile}, nil
}

// SetDeliveredStatus implements the handler interface.
func (s *handler) SetDeliveredStatus(ctx context.Context, req *domain.SetDeliveredStatusRequest) (resp *domain.EmptyResponse, err error) {
	err = stockpileService.SetDeliveredStatus(ctx, req.ProductId, req.Amount, req.Status)
	if err != nil {
		return nil, err
	}
	return &domain.EmptyResponse{}, nil
}
