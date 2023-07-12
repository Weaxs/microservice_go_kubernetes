package api

import (
	"context"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
)

var PaymentApiServiceInfo = newPaymentApiServiceInfo()

func newPaymentApiServiceInfo() *kitex.ServiceInfo {
	serviceName := "PaymentApi"
	handlerType := (*PaymentApi)(nil)
	methods := map[string]kitex.MethodInfo{
		"executeSettlement":       kitex.NewMethodInfo(executeSettlementHandler, newExecuteSettlementArgs, newExecuteSettlementResult, false),
		"updatePaymentState":      kitex.NewMethodInfo(updatePaymentStateHandler, newUpdatePaymentStateArgs, newUpdatePaymentStateResult, false),
		"updatePaymentStateAlias": kitex.NewMethodInfo(updatePaymentStateAliasHandler, newUpdatePaymentStateAliasArgs, newUpdatePaymentStateAliasResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "payment",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func executeSettlementHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(domain.ExecuteSettlementRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(PaymentApi).ExecuteSettlement(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *domain.ExecuteSettlementArgs:
		success, err := handler.(PaymentApi).ExecuteSettlement(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*domain.ExecuteSettlementResult)
		realResult.Success = success
	}
	return nil
}

func updatePaymentStateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(domain.UpdatePaymentStateRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(PaymentApi).UpdatePaymentState(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *domain.UpdatePaymentStateArgs:
		success, err := handler.(PaymentApi).UpdatePaymentState(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*domain.UpdatePaymentStateResult)
		realResult.Success = success
	}
	return nil
}

func updatePaymentStateAliasHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(domain.UpdatePaymentStateAlias)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(PaymentApi).UpdatePaymentStateAlias(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *domain.UpdatePaymentStateAliasArgs:
		success, err := handler.(PaymentApi).UpdatePaymentStateAlias(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*domain.UpdatePaymentStateAliasResult)
		realResult.Success = success
	}
	return nil
}

func newUpdatePaymentStateArgs() interface{} {
	return &domain.UpdatePaymentStateArgs{}
}

func newUpdatePaymentStateResult() interface{} {
	return &domain.UpdatePaymentStateResult{}
}

func newExecuteSettlementArgs() interface{} {
	return &domain.ExecuteSettlementArgs{}
}

func newExecuteSettlementResult() interface{} {
	return &domain.ExecuteSettlementResult{}
}

func newUpdatePaymentStateAliasArgs() interface{} {
	return &domain.UpdatePaymentStateAliasArgs{}
}

func newUpdatePaymentStateAliasResult() interface{} {
	return &domain.UpdatePaymentStateAliasResult{}
}
