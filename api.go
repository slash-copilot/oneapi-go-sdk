package oneapigosdk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

var (
	PATH_TOKEN          = "/api/token"
	PATH_RECHARGE_TOKEN = "/api/token/recharge"
	PATH_TOKEN_INFO     = "/api/token/info"
	PATH_USER_LOGS      = "/api/log/self"
)

type Token struct {
	Id             int    `json:"id"`
	UserId         int    `json:"user_id"`
	Key            string `json:"key"`
	Status         int    `json:"status"`
	Name           string `json:"name"`
	CreatedTime    int64  `json:"created_time"`
	AccessedTime   int64  `json:"accessed_time"`
	ExpiredTime    int64  `json:"expired_time"`
	RemainQuota    int    `json:"remain_quota"`
	UnlimitedQuota bool   `json:"unlimited_quota"`
	UsedQuota      int    `json:"used_quota"`
}

type AddTokenReq struct {
	Name           string `json:"name"`
	Remark 	   	   string `json:"remark"`
	ExpiredTime    int64  `json:"expired_time"`
	RemainQuota    int    `json:"remain_quota"`
	UnlimitedQuota bool   `json:"unlimited_quota"`
}

type AddTokenResp struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Data    *Token `json:"data"`
}

type UpdateTokenReq struct {
	Key            string `json:"key"`
	Name           string `json:"name"`
	ExpiredTime    int64  `json:"expired_time"`
	RemainQuota    int    `json:"remain_quota"`
	UnlimitedQuota bool   `json:"unlimited_quota"`
}

type RechargeTokenReq struct {
	Key    string `json:"key"`
	Amount int    `json:"amount"`
}

type RechargeTokenResp struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type GetTokenResp struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Data    *Token `json:"data"`
}

type UpdateTokenResp struct {
	Data    *Token `json:"data"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type Log struct {
	Id               int    `json:"id"`
	UserId           int    `json:"user_id"`
	CreatedAt        int64  `json:"created_at"`
	Type             int    `json:"type"`
	Content          string `json:"content"`
	Username         string `json:"username"`
	TokenName        string `json:"token_name"`
	ModelName        string `json:"model_name"`
	Quota            int    `json:"quota"`
	PromptTokens     int    `json:"prompt_tokens"`
	CompletionTokens int    `json:"completion_tokens"`
	ChannelId        int    `json:"channel"`
}

type GetUserLogsReq struct {
	P              int     `json:"p"`
	Type           []int   `json:"type"`
	TokenName      string  `json:"token_name"`
	ModelName      *string `json:"model_name"`
	StartTimestamp *int64  `json:"start_timestamp"`
	EndTimestamp   *int64  `json:"end_timestamp"`
}

type GetUserLogsResp struct {
	Data    []*Log `json:"data"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type Api struct {
	c *Client
}

func (api *Api) buildRequestApi(requestUrl string) string {
	return api.c.GetHost() + requestUrl
}

func (api *Api) createBaseRequest(ctx context.Context, method string, url string, req ...interface{}) (r *http.Request, err error) {
	if r, err = api.c.NewHttpClient(ctx, method, url, req...); err == nil {
		api.c.SetHttpRequest(r).
			SetHttpRequestHeader("Authorization", "Bearer "+api.c.GetAccessToken()).
			SetHttpRequestHeader("Cache-Control", "no-cache").
			SetHttpRequestHeader("Content-Type", "application/json; text/plain")
	}
	return
}

func (api *Api) createGetRequest(ctx context.Context, url string, req ...interface{}) (*http.Request, error) {
	return api.createBaseRequest(ctx, http.MethodGet, url, req...)
}

func (api *Api) createPostRequest(ctx context.Context, url string, req ...interface{}) (*http.Request, error) {
	return api.createBaseRequest(ctx, http.MethodPost, url, req...)
}

func (api *Api) createPatchRequest(ctx context.Context, url string, req ...interface{}) (*http.Request, error) {
	return api.createBaseRequest(ctx, http.MethodPatch, url, req...)
}

func (api *Api) AddToken(ctx context.Context, req *AddTokenReq) (resp *AddTokenResp, err error) {
	var r *http.Request
	if r, err = api.createPostRequest(ctx, api.buildRequestApi(PATH_TOKEN), req); err != nil {
		return
	}

	var _resp AddTokenResp

	err = api.c.SetHttpRequest(r).SendRequest(&_resp)
	resp = &_resp
	return
}

func (api *Api) UpdateToken(ctx context.Context, req *UpdateTokenReq) (resp *UpdateTokenResp, err error) {
	var r *http.Request
	if r, err = api.createPatchRequest(ctx, api.buildRequestApi(PATH_TOKEN), req); err != nil {
		return
	}

	var _resp UpdateTokenResp

	err = api.c.SetHttpRequest(r).SendRequest(&_resp)
	resp = &_resp
	return
}

func (api *Api) RechargeToken(ctx context.Context, req *RechargeTokenReq) (resp *RechargeTokenResp, err error) {
	var r *http.Request
	if r, err = api.createPostRequest(ctx, api.buildRequestApi(PATH_RECHARGE_TOKEN), req); err != nil {
		return
	}

	var _resp RechargeTokenResp

	err = api.c.SetHttpRequest(r).SendRequest(&_resp)
	resp = &_resp
	return
}

func (api *Api) GetTokenStatus(ctx context.Context, key string) (resp *GetTokenResp, err error) {
	var r *http.Request
	if r, err = api.createGetRequest(ctx, api.buildRequestApi(PATH_TOKEN_INFO+"/"+key)); err != nil {
		return
	}

	var _resp GetTokenResp
	err = api.c.SetHttpRequest(r).SendRequest(&_resp)
	resp = &_resp
	return
}

func (api *Api) GetUserLogs(ctx context.Context, req *GetUserLogsReq) (resp *GetUserLogsResp, err error) {
	var u = url.Values{}
	u.Set("p", fmt.Sprintf("%d", req.P))
	u.Set("token_name", req.TokenName)

	if req.Type != nil {
		for _, v := range req.Type {
			u.Add("type", fmt.Sprintf("%d", v))
		}
	}

	if req.StartTimestamp != nil {
		u.Set("start_timestamp", fmt.Sprintf("%d", *req.StartTimestamp))
	}
	if req.EndTimestamp != nil {
		u.Set("end_timestamp", fmt.Sprintf("%d", *req.EndTimestamp))
	}

	var r *http.Request
	if r, err = api.createGetRequest(ctx, api.buildRequestApi(PATH_USER_LOGS), u); err != nil {
		return
	}

	var _resp GetUserLogsResp
	err = api.c.SetHttpRequest(r).SendRequest(&_resp)
	resp = &_resp
	return
}
