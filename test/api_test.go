package test

import (
	"context"
	"encoding/json"
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

func TestAddToken(t *testing.T) {
	var client = oneapigosdk.NewClient(host, accessToken)
	var err error
	var ctx, _ = context.WithCancel(context.Background())

	var res *oneapigosdk.AddTokenResp
	if res, err = client.Api().AddToken(ctx, &oneapigosdk.Token{
		Name:           "test",
		RemainQuota:    0,
		ExpiredTime:    -1,
		UnlimitedQuota: false,
	}); err != nil {
		t.Fatal(err.Error())
	}

	j, _ := json.Marshal(res)

	t.Log(string(j))
}
