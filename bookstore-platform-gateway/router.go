package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/route"
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

	})
	h.POST("/accounts", func(c context.Context, ctx *app.RequestContext) {

	})
	h.PUT("/account", func(c context.Context, ctx *app.RequestContext) {

	})
}

func registerAuthorizationRoute(h *route.RouterGroup) {
	h.GET("/token", TokenHandler)
}

func registerPaymentRoute(h *route.RouterGroup) {
	h.POST("/settlements", func(c context.Context, ctx *app.RequestContext) {

	})
	h.PATCH("/pay/:id?state=PAYED", func(c context.Context, ctx *app.RequestContext) {

	})
	h.PATCH("/pay/:id?state=CANCEL", func(c context.Context, ctx *app.RequestContext) {

	})
}

func registerWarehouseRoute(h *route.RouterGroup) {
	h.GET("/products", func(c context.Context, ctx *app.RequestContext) {

	})
	h.GET("/products/:id", func(c context.Context, ctx *app.RequestContext) {

	})
	h.POST("/products", func(c context.Context, ctx *app.RequestContext) {

	})
	h.PUT("/products", func(c context.Context, ctx *app.RequestContext) {

	})
	h.DELETE("products/:productId", func(c context.Context, ctx *app.RequestContext) {

	})
	h.GET("/products/stockpile/:productId", func(c context.Context, ctx *app.RequestContext) {

	})
	h.PATCH("/products/stockpile/:productId?amount=:amount", func(c context.Context, ctx *app.RequestContext) {

	})
	h.GET("/advertisements", func(c context.Context, ctx *app.RequestContext) {

	})

}
