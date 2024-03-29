package test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	oneapigosdk "github.com/slash-copilot/oneapi-go-sdk"
)

var (
	ONEAPI_HOST         = os.Getenv("ONEAPI_HOST")
	ONEAPI_ACCESS_TOKEN = os.Getenv("ONEAPI_ACCESS_TOKEN")
	ONEAPI_API_KEY      = os.Getenv("ONEAPI_API_KEY")
)

func TestClient(t *testing.T) {
	var c = &oneapigosdk.ClientConfig{
		Host:        ONEAPI_HOST,
		AccessToken: ONEAPI_ACCESS_TOKEN,
	}

	var client1 = oneapigosdk.NewClientWithConfig(c)

	var client2 = oneapigosdk.NewClient(ONEAPI_HOST, ONEAPI_ACCESS_TOKEN)

	t.Log(client1.GetHost() == client2.GetHost())
	t.Log(client1.GetAccessToken() == client2.GetAccessToken())
}

func TestAddToken(t *testing.T) {
	var client = oneapigosdk.NewClient(ONEAPI_HOST, ONEAPI_ACCESS_TOKEN)
	var err error
	var ctx = context.Background()

	var res *oneapigosdk.AddTokenResp
	if res, err = client.Api().AddToken(ctx, &oneapigosdk.AddTokenReq{
		Name:           "test2",
		RemainQuota:    0,
		ExpiredTime:    -1,
		UnlimitedQuota: false,
	}); err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("res: %v", res)
}

func TestUpdateToken(t *testing.T) {
	var client = oneapigosdk.NewClient(ONEAPI_HOST, ONEAPI_ACCESS_TOKEN)
	var err error
	var ctx = context.Background()

	var res *oneapigosdk.UpdateTokenResp
	if res, err = client.Api().UpdateToken(ctx, &oneapigosdk.UpdateTokenReq{
		Key:            ONEAPI_API_KEY,
		Name:           "test2333",
		RemainQuota:    10,
		ExpiredTime:    -1,
		UnlimitedQuota: false,
	}); err != nil {
		t.Fatal(err.Error())
	}
	t.Log(res)
}

func TestGetUserLogs(t *testing.T) {
	var client = oneapigosdk.NewClient(ONEAPI_HOST, ONEAPI_ACCESS_TOKEN)
	var err error
	var ctx = context.Background()

	var res *oneapigosdk.GetUserLogsResp

	if res, err = client.Api().GetUserLogs(ctx, &oneapigosdk.GetUserLogsReq{
		P:    1,
		Type: []int{1, 2},
	}); err != nil {
		t.Fatal(err.Error())
	}
	j, _ := json.Marshal(res)
	t.Log(string(j))
}

func TestRechargeToken(t *testing.T) {
	var client = oneapigosdk.NewClient(ONEAPI_HOST, ONEAPI_ACCESS_TOKEN)
	var err error
	var ctx = context.Background()

	var res *oneapigosdk.RechargeTokenResp

	if res, err = client.Api().RechargeToken(ctx, &oneapigosdk.RechargeTokenReq{
		Key:    ONEAPI_API_KEY,
		Amount: 1000,
	}); err != nil {
		t.Fatal(err.Error())
	}
	j, _ := json.Marshal(res)
	t.Log(string(j))
}
