package client

import (
	"context"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/api"
	warehouse "github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// AccountClient is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type WarehouseClient interface {
	GetAllAdvertisements(ctx context.Context, Req *warehouse.Empty, callOptions ...callopt.Option) (r *warehouse.GetAllAdvertisementsResponse, err error)
	GetAllProducts(ctx context.Context, Req *warehouse.Empty, callOptions ...callopt.Option) (r *warehouse.GetAllProductResponse, err error)
	GetProduct(ctx context.Context, Req *warehouse.GetProductRequest, callOptions ...callopt.Option) (r *warehouse.GetProductResponse, err error)
	CreateProduct(ctx context.Context, Req *warehouse.ChangeProductRequest, callOptions ...callopt.Option) (r *warehouse.Empty, err error)
	UpdateProduct(ctx context.Context, Req *warehouse.ChangeProductRequest, callOptions ...callopt.Option) (r *warehouse.Empty, err error)
	RemoveProduct(ctx context.Context, Req *warehouse.RemoveProductRequest, callOptions ...callopt.Option) (r *warehouse.Empty, err error)
	UpdateStockpile(ctx context.Context, Req *warehouse.UpdateStockpileRequest, callOptions ...callopt.Option) (r *warehouse.Empty, err error)
	QueryStockpile(ctx context.Context, Req *warehouse.QueryStockpileRequest, callOptions ...callopt.Option) (r *warehouse.QueryStockpileResponse, err error)
	SetDeliveredStatus(ctx context.Context, Req *warehouse.SetDeliveredStatusRequest, callOptions ...callopt.Option) (r *warehouse.Empty, err error)
}

// NewAccountClient creates a client for the service defined in IDL.
func NewWarehouseClient(opts ...client.Option) (WarehouseClient, error) {
	var options []client.Option
	options = append(options, client.WithDestService("WarehouseApi"))

	options = append(options, opts...)

	kc, err := client.NewClient(api.WarehouseApiServiceInfo, options...)
	if err != nil {
		return nil, err
	}
	return &kWarehouseApiClient{
		kClient: newServiceClient(kc),
	}, nil
}

type kWarehouseApiClient struct {
	*kClient
}

func (p *kWarehouseApiClient) GetAllAdvertisements(ctx context.Context, Req *warehouse.Empty, callOptions ...callopt.Option) (r *warehouse.GetAllAdvertisementsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAllAdvertisements(ctx, Req)
}

func (p *kWarehouseApiClient) GetAllProducts(ctx context.Context, Req *warehouse.Empty, callOptions ...callopt.Option) (r *warehouse.GetAllProductResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAllProducts(ctx, Req)
}

func (p *kWarehouseApiClient) GetProduct(ctx context.Context, Req *warehouse.GetProductRequest, callOptions ...callopt.Option) (r *warehouse.GetProductResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProduct(ctx, Req)
}

func (p *kWarehouseApiClient) CreateProduct(ctx context.Context, Req *warehouse.ChangeProductRequest, callOptions ...callopt.Option) (r *warehouse.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateProduct(ctx, Req)
}

func (p *kWarehouseApiClient) UpdateProduct(ctx context.Context, Req *warehouse.ChangeProductRequest, callOptions ...callopt.Option) (r *warehouse.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateProduct(ctx, Req)
}

func (p *kWarehouseApiClient) RemoveProduct(ctx context.Context, Req *warehouse.RemoveProductRequest, callOptions ...callopt.Option) (r *warehouse.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RemoveProduct(ctx, Req)
}

func (p *kWarehouseApiClient) UpdateStockpile(ctx context.Context, Req *warehouse.UpdateStockpileRequest, callOptions ...callopt.Option) (r *warehouse.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateStockpile(ctx, Req)
}

func (p *kWarehouseApiClient) QueryStockpile(ctx context.Context, Req *warehouse.QueryStockpileRequest, callOptions ...callopt.Option) (r *warehouse.QueryStockpileResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryStockpile(ctx, Req)
}

func (p *kWarehouseApiClient) SetDeliveredStatus(ctx context.Context, Req *warehouse.SetDeliveredStatusRequest, callOptions ...callopt.Option) (r *warehouse.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetDeliveredStatus(ctx, Req)
}

func (p *kClient) GetAllAdvertisements(ctx context.Context, Req *warehouse.Empty) (r *warehouse.GetAllAdvertisementsResponse, err error) {
	var _args warehouse.GetAllAdvertisementsArgs
	_args.Req = Req
	var _result warehouse.GetAllAdvertisementsResult
	if err = p.c.Call(ctx, "GetAllAdvertisements", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAllProducts(ctx context.Context, Req *warehouse.Empty) (r *warehouse.GetAllProductResponse, err error) {
	var _args warehouse.GetAllProductsArgs
	_args.Req = Req
	var _result warehouse.GetAllProductsResult
	if err = p.c.Call(ctx, "GetAllProducts", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetProduct(ctx context.Context, Req *warehouse.GetProductRequest) (r *warehouse.GetProductResponse, err error) {
	var _args warehouse.GetProductArgs
	_args.Req = Req
	var _result warehouse.GetProductResult
	if err = p.c.Call(ctx, "GetProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateProduct(ctx context.Context, Req *warehouse.ChangeProductRequest) (r *warehouse.Empty, err error) {
	var _args warehouse.CreateProductArgs
	_args.Req = Req
	var _result warehouse.CreateProductResult
	if err = p.c.Call(ctx, "CreateProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateProduct(ctx context.Context, Req *warehouse.ChangeProductRequest) (r *warehouse.Empty, err error) {
	var _args warehouse.UpdateProductArgs
	_args.Req = Req
	var _result warehouse.UpdateProductResult
	if err = p.c.Call(ctx, "UpdateProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RemoveProduct(ctx context.Context, Req *warehouse.RemoveProductRequest) (r *warehouse.Empty, err error) {
	var _args warehouse.RemoveProductArgs
	_args.Req = Req
	var _result warehouse.RemoveProductResult
	if err = p.c.Call(ctx, "RemoveProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateStockpile(ctx context.Context, Req *warehouse.UpdateStockpileRequest) (r *warehouse.Empty, err error) {
	var _args warehouse.UpdateStockpileArgs
	_args.Req = Req
	var _result warehouse.UpdateStockpileResult
	if err = p.c.Call(ctx, "UpdateStockpile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryStockpile(ctx context.Context, Req *warehouse.QueryStockpileRequest) (r *warehouse.QueryStockpileResponse, err error) {
	var _args warehouse.QueryStockpileArgs
	_args.Req = Req
	var _result warehouse.QueryStockpileResult
	if err = p.c.Call(ctx, "QueryStockpile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SetDeliveredStatus(ctx context.Context, Req *warehouse.SetDeliveredStatusRequest) (r *warehouse.Empty, err error) {
	var _args warehouse.SetDeliveredStatusArgs
	_args.Req = Req
	var _result warehouse.SetDeliveredStatusResult
	if err = p.c.Call(ctx, "SetDeliveredStatus", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
