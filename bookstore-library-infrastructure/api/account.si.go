package api

import (
	"context"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

var AccountApiServiceInfo = newAccountApiServiceInfo()

func newAccountApiServiceInfo() *serviceinfo.ServiceInfo {
	serviceName := "AccountApi"
	handlerType := (*AccountApi)(nil)
	methods := map[string]serviceinfo.MethodInfo{
		"getAccount":    serviceinfo.NewMethodInfo(getAccountHandler, newAccountApiGetAccountArgs, newAccountApiGetAccountResult, false),
		"createAccount": serviceinfo.NewMethodInfo(createAccountHandler, newAccountChangeAccountArgs, newAccountChangeAccountResult, false),
		"updateAccount": serviceinfo.NewMethodInfo(updateAccountHandler, newAccountChangeAccountArgs, newAccountChangeAccountResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "account",
	}
	svcInfo := &serviceinfo.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    serviceinfo.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func getAccountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*domain.GetAccountArgs)
	realResult := result.(*domain.GetAccountResult)
	success, err := handler.(AccountApi).GetAccount(ctx, realArg.Username)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAccountApiGetAccountArgs() interface{} {
	return domain.NewGetAccountArgs()
}

func newAccountApiGetAccountResult() interface{} {
	return domain.NewGetAccountResult()
}

func createAccountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*domain.ChangeAccountArgs)
	realResult := result.(*domain.ChangeAccountResult)
	success, err := handler.(AccountApi).CreateAccount(ctx, realArg.Account)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAccountChangeAccountArgs() interface{} {
	return domain.NewAccountChangeAccountArgs()
}

func newAccountChangeAccountResult() interface{} {
	return domain.NewAccountChangeAccountResult()
}

func updateAccountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*domain.ChangeAccountArgs)
	realResult := result.(*domain.ChangeAccountResult)
	success, err := handler.(AccountApi).UpdateAccount(ctx, realArg.Account)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
