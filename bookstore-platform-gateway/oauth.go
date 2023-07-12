package main

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"golang.org/x/oauth2"
)

func TokenHandler(c context.Context, ctx *app.RequestContext) {
	username := ctx.Param("username")

	var t *oauth2.Token
	var err error
	if username != "" {
		t, err = token(c, ctx)
	} else {
		t, err = refresh(c, ctx)
	}
	b, err := json.Marshal(t)
	if err != nil {

	}
	ctx.JSON(consts.StatusOK, string(b))
}

func token(c context.Context, ctx *app.RequestContext) (*oauth2.Token, error) {

	username := ctx.Param("username")
	password := ctx.Param("password")
	grantType := ctx.Param("grant_type")
	clientId := ctx.Param("client_id")
	clientSecret := ctx.Param("client_secret")
	conf := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{grantType},
	}
	return conf.PasswordCredentialsToken(c, username, password)
}

func refresh(c context.Context, ctx *app.RequestContext) (*oauth2.Token, error) {
	refreshToken := ctx.Param("refresh_token")
	grantType := ctx.Param("grant_type")
	clientId := ctx.Param("client_id")
	clientSecret := ctx.Param("client_secret")

	t := &oauth2.Token{RefreshToken: refreshToken}
	conf := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{grantType},
	}
	return conf.TokenSource(c, t).Token()
}
