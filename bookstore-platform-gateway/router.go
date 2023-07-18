package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/hertz/pkg/route"
)

var (
	account   = NewAccountService()
	warehouse = NewWarehouseService()
	payment   = NewPaymentService()
)

func RegisterRoute(h *server.Hertz) {
	restful := h.Group("/restful")
	registerAccountRoute(restful)
	registerPaymentRoute(restful)
	registerWarehouseRoute(restful)

	oauth := h.Group("/oauth")
	registerAuthorizationRoute(oauth)

}

func registerAccountRoute(h *route.RouterGroup) {
	h.GET("/accounts/:username", func(c context.Context, ctx *app.RequestContext) {
		account.findAccountByUsername(c, ctx)
	})
	h.POST("/accounts", func(c context.Context, ctx *app.RequestContext) {
		account.createAccount(c, ctx)
	})
	h.PUT("/account", func(c context.Context, ctx *app.RequestContext) {
		account.updateAccount(c, ctx)
	})
}

func registerAuthorizationRoute(h *route.RouterGroup) {
	h.GET("/token", TokenHandler)
}

func registerPaymentRoute(h *route.RouterGroup) {
	h.POST("/settlements", func(c context.Context, ctx *app.RequestContext) {
		payment.executeSettlement(c, ctx)
	})
	h.PATCH("/pay/:id", func(c context.Context, ctx *app.RequestContext) {
		state := ctx.Param("state")
		if state == "PAYED" || state == "CANCEL" {
			payment.modifyPaymentState(c, ctx)
		} else {
			ctx.JSON(consts.StatusBadRequest, map[string]string{"msg": "State parameter is empty."})
		}
	})
}

func registerWarehouseRoute(h *route.RouterGroup) {
	h.GET("/products", func(c context.Context, ctx *app.RequestContext) {
		warehouse.getAllProducts(c, ctx)
	})
	h.GET("/products/:id", func(c context.Context, ctx *app.RequestContext) {
		warehouse.getProductById(c, ctx)
	})
	h.POST("/products", func(c context.Context, ctx *app.RequestContext) {
		warehouse.createProduct(c, ctx)
	})
	h.PUT("/products", func(c context.Context, ctx *app.RequestContext) {
		warehouse.updateProduct(c, ctx)
	})
	h.DELETE("products/:productId", func(c context.Context, ctx *app.RequestContext) {
		warehouse.deleteById(c, ctx)
	})
	h.GET("/products/stockpile/:productId", func(c context.Context, ctx *app.RequestContext) {
		warehouse.getStockpileByProductId(c, ctx)
	})
	h.PATCH("/products/stockpile/:productId?amount=:amount", func(c context.Context, ctx *app.RequestContext) {
		warehouse.updateStockpile(c, ctx)
	})
	h.GET("/advertisements", func(c context.Context, ctx *app.RequestContext) {

	})

}
