package main

import (
	"context"
	warehouse "github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
)

// handler implements the last service interface defined in the IDL.
type handler struct{}

// GetAllAdvertisements implements the handler interface.
func (s *handler) GetAllAdvertisements(ctx context.Context, req *warehouse.EmptyResponse) (resp *warehouse.GetAllAdvertisementsResponse, err error) {
	// TODO: Your code here...
	return
}

// GetAllProducts implements the handler interface.
func (s *handler) GetAllProducts(ctx context.Context, req *warehouse.EmptyResponse) (resp *warehouse.GetAllProductResponse, err error) {
	// TODO: Your code here...
	return
}

// GetProduct implements the handler interface.
func (s *handler) GetProduct(ctx context.Context, req *warehouse.GetProductRequest) (resp *warehouse.GetProductResponse, err error) {
	// TODO: Your code here...
	return
}

// CreateProduct implements the handler interface.
func (s *handler) CreateProduct(ctx context.Context, req *warehouse.ChangeProductRequest) (resp *warehouse.EmptyResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateProduct implements the handler interface.
func (s *handler) UpdateProduct(ctx context.Context, req *warehouse.ChangeProductRequest) (resp *warehouse.EmptyResponse, err error) {
	// TODO: Your code here...
	return
}

// RemoveProduct implements the handler interface.
func (s *handler) RemoveProduct(ctx context.Context, req *warehouse.RemoveProductRequest) (resp *warehouse.EmptyResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateStockpile implements the handler interface.
func (s *handler) UpdateStockpile(ctx context.Context, req *warehouse.UpdateStockpileRequest) (resp *warehouse.EmptyResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryStockpile implements the handler interface.
func (s *handler) QueryStockpile(ctx context.Context, req *warehouse.QueryStockpileRequest) (resp *warehouse.QueryStockpileResponse, err error) {
	// TODO: Your code here...
	return
}

// SetDeliveredStatus implements the handler interface.
func (s *handler) SetDeliveredStatus(ctx context.Context, req *warehouse.SetDeliveredStatusRequest) (resp *warehouse.EmptyResponse, err error) {
	// TODO: Your code here...
	return
}
