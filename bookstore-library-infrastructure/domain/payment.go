package domain

import (
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 支付状态
type PaymentState int32

const (
	PaymentState_WAITING     PaymentState = 0
	PaymentState_CANCEL      PaymentState = 1
	PaymentState_PAYED       PaymentState = 2
	PaymentState_TIMEOUT     PaymentState = 3
	PaymentState_NOT_SUPPORT PaymentState = 4
)

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount    int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	ProductId int64 `protobuf:"varint,2,opt,name=productId,proto3" json:"productId,omitempty"`
}

type Purchase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Delivery  bool   `protobuf:"varint,1,opt,name=delivery,proto3" json:"delivery,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Telephone string `protobuf:"bytes,3,opt,name=telephone,proto3" json:"telephone,omitempty"`
	Location  string `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
}

type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateTime  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=createTime,proto3" json:"createTime,omitempty"`
	PayId       string                 `protobuf:"bytes,2,opt,name=payId,proto3" json:"payId,omitempty"`
	TotalPrice  float64                `protobuf:"fixed64,3,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
	Expires     int64                  `protobuf:"varint,4,opt,name=expires,proto3" json:"expires,omitempty"`
	PaymentLink string                 `protobuf:"bytes,5,opt,name=paymentLink,proto3" json:"paymentLink,omitempty"`
	PayState    PaymentState           `protobuf:"varint,6,opt,name=payState,proto3,enum=payment.PaymentState" json:"payState,omitempty"`
}

// 支付结算单
type Settlement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items      []*Item            `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Purchase   *Purchase          `protobuf:"bytes,2,opt,name=purchase,proto3" json:"purchase,omitempty"`
	ProductMap map[int64]*Product `protobuf:"bytes,3,rep,name=productMap,proto3" json:"productMap,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

type ExecuteSettlementRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Settlement *Settlement `protobuf:"bytes,1,opt,name=settlement,proto3" json:"settlement,omitempty"`
}

type ExecuteSettlementResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payment *Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
}

type UpdatePaymentStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayId string       `protobuf:"bytes,1,opt,name=payId,proto3" json:"payId,omitempty"`
	State PaymentState `protobuf:"varint,2,opt,name=state,proto3,enum=payment.PaymentState" json:"state,omitempty"`
}

type UpdatePaymentStateAlias struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayId     string       `protobuf:"bytes,1,opt,name=payId,proto3" json:"payId,omitempty"`
	AccountId int64        `protobuf:"varint,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
	State     PaymentState `protobuf:"varint,3,opt,name=state,proto3,enum=payment.PaymentState" json:"state,omitempty"`
}

type ExecuteSettlementArgs struct {
	Req *ExecuteSettlementRequest
}

type ExecuteSettlementResult struct {
	Success *ExecuteSettlementResponse
}

type UpdatePaymentStateArgs struct {
	Req *UpdatePaymentStateRequest
}

type UpdatePaymentStateResult struct {
	Success *EmptyResponse
}

type UpdatePaymentStateAliasArgs struct {
	Req *UpdatePaymentStateAlias
}

type UpdatePaymentStateAliasResult struct {
	Success *EmptyResponse
}
