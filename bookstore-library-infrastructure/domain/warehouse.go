package domain

import "google.golang.org/protobuf/runtime/protoimpl"

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// DeliveredStatus 待交付商品的状态
type DeliveredStatus int32

const (
	DeliveredStatus_DECREASE DeliveredStatus = 0
	DeliveredStatus_INCREASE DeliveredStatus = 1
	DeliveredStatus_FROZEN   DeliveredStatus = 2
	DeliveredStatus_THAWED   DeliveredStatus = 3
)

// Enum value maps for DeliveredStatus.
var (
	DeliveredStatus_name = map[int32]string{
		0: "DECREASE",
		1: "INCREASE",
		2: "FROZEN",
		3: "THAWED",
	}
	DeliveredStatus_value = map[string]int32{
		"DECREASE": 0,
		"INCREASE": 1,
		"FROZEN":   2,
		"THAWED":   3,
	}
)

// Advertisement 广告对象模型
type Advertisement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image     string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	ProductId int64  `protobuf:"varint,2,opt,name=productId,proto3" json:"productId,omitempty"`
}

// Specification 商品规格
type Specification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item      string `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Value     string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ProductId int64  `protobuf:"varint,3,opt,name=productId,proto3" json:"productId,omitempty"`
}

// Product 商品对象
type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title          string           `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Price          float64          `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	Rate           float32          `protobuf:"fixed32,3,opt,name=rate,proto3" json:"rate,omitempty"`
	Description    string           `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Cover          string           `protobuf:"bytes,5,opt,name=cover,proto3" json:"cover,omitempty"`
	Detail         string           `protobuf:"bytes,6,opt,name=detail,proto3" json:"detail,omitempty"`
	Specifications []*Specification `protobuf:"bytes,7,rep,name=specifications,proto3" json:"specifications,omitempty"`
}

// Stockpile 商品库存
type Stockpile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount  int64    `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Frozen  int64    `protobuf:"varint,2,opt,name=frozen,proto3" json:"frozen,omitempty"`
	Product *Product `protobuf:"bytes,3,opt,name=product,proto3" json:"product,omitempty"`
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

type GetAllAdvertisementsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Advertisements []*Advertisement `protobuf:"bytes,1,rep,name=advertisements,proto3" json:"advertisements,omitempty"`
}

type GetAllProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

type GetProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

type RemoveProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

type GetProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

type ChangeProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

type UpdateStockpileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int64 `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Amount    int64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

type QueryStockpileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int64 `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
}

type QueryStockpileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stockpile *Stockpile `protobuf:"bytes,1,opt,name=stockpile,proto3" json:"stockpile,omitempty"`
}

type SetDeliveredStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int64           `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Amount    int64           `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Status    DeliveredStatus `protobuf:"varint,3,opt,name=status,proto3,enum=warehouse.DeliveredStatus" json:"status,omitempty"`
}

type QueryStockpileResult struct {
	Success *QueryStockpileResponse
}

type GetAllAdvertisementsArgs struct {
	Req *EmptyResponse
}

type GetAllAdvertisementsResult struct {
	Success *GetAllAdvertisementsResponse
}

type GetAllProductsResult struct {
	Success *GetAllProductResponse
}

type GetAllProductsArgs struct {
	Req *EmptyResponse
}

type GetProductArgs struct {
	Req *GetProductRequest
}

type GetProductResult struct {
	Success *GetProductResponse
}

type CreateProductArgs struct {
	Req *ChangeProductRequest
}

type CreateProductResult struct {
	Success *EmptyResponse
}

type UpdateProductArgs struct {
	Req *ChangeProductRequest
}

type QueryStockpileArgs struct {
	Req *QueryStockpileRequest
}

type SetDeliveredStatusArgs struct {
	Req *SetDeliveredStatusRequest
}
type SetDeliveredStatusResult struct {
	Success *EmptyResponse
}
type RemoveProductArgs struct {
	Req *RemoveProductRequest
}

type RemoveProductResult struct {
	Success *EmptyResponse
}
type UpdateStockpileArgs struct {
	Req *UpdateStockpileRequest
}
type UpdateStockpileResult struct {
	Success *EmptyResponse
}

type UpdateProductResult struct {
	Success *EmptyResponse
}
