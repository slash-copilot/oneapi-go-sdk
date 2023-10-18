package test

import (
	"testing"

	oneapigosdk "github.com/slash-copilot/oneapi-go-sdk"
)

var (
	host        = "host"
	accessToken = "accessToken"
)

func TestClient(t *testing.T) {
	var c = &oneapigosdk.ClientConfig{
		Host:        host,
		AccessToken: accessToken,
	}

	var client1 = oneapigosdk.NewClientWithConfig(c)

	var client2 = oneapigosdk.NewClient(host, accessToken)

	t.Log(client1.GetHost() == client2.GetHost())
	t.Log(client1.GetAccessToken() == client2.GetAccessToken())
}
