package client

import (
	"context"
	"github.com/Weaxs/microservice_go_kubernetes/bookstore-library-infrastructure/domain"
	"github.com/cloudwego/kitex/client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountClient(t *testing.T) {
	client, err := NewAccountClient(
		client.WithHostPorts("[::1]:8810"),
		client.WithMuxConnection(1))
	if err != nil {
		panic(any(err))
	}
	account, err := client.GetAccount(context.Background(), &domain.GetAccountRequest{Username: "icyfenix"})
	if err != nil {
		panic(any(err))
	}
	assert.Equal(t, "icyfenix", account.Account.Username)
}
