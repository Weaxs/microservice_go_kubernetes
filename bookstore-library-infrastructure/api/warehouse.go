package api

import (
	"context"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
)

type WarehouseApi interface {
	GetAllAdvertisements(ctx context.Context, req *domain.Empty) (res *domain.GetAllAdvertisementsResponse, err error)
	GetAllProducts(ctx context.Context, req *domain.Empty) (res *domain.GetAllProductResponse, err error)
	GetProduct(ctx context.Context, req *domain.GetProductRequest) (res *domain.GetProductResponse, err error)
	CreateProduct(ctx context.Context, req *domain.ChangeProductRequest) (res *domain.Empty, err error)
	UpdateProduct(ctx context.Context, req *domain.ChangeProductRequest) (res *domain.Empty, err error)
	RemoveProduct(ctx context.Context, req *domain.RemoveProductRequest) (res *domain.Empty, err error)
	UpdateStockpile(ctx context.Context, req *domain.UpdateStockpileRequest) (res *domain.Empty, err error)
	QueryStockpile(ctx context.Context, req *domain.QueryStockpileRequest) (res *domain.QueryStockpileResponse, err error)
	SetDeliveredStatus(ctx context.Context, req *domain.SetDeliveredStatusRequest) (res *domain.Empty, err error)
}
