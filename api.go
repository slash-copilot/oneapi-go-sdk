package oneapigosdk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

var (
	PATH_TOKEN      = "/api/token"
	PATH_TOKEN_INFO = "/api/token/info"
	PATH_ALL_LOGS   = "/api/log"
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

type GetAllLogsReq struct {
	P              *int    `json:"p"`
	Type           *int    `json:"type"`	
	TokenName      *string `json:"token_name"`
	ModelName      *string `json:"model_name"`
	StartTimestamp *int64  `json:"start_timestamp"`
	EndTimestamp   *int64  `json:"end_timestamp"`
	Channel        *int    `json:"channel"`
}

type GetAllLogsResp struct {
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

func (api *Api) GetAllLogs(ctx context.Context, req *GetAllLogsReq) (resp *GetAllLogsResp, err error) {
	var u = url.Values{}
	if req.P != nil {
		u.Set("p", fmt.Sprintf("%d", *req.P))
	}
	if req.Type != nil {
		u.Set("type", fmt.Sprintf("%d", *req.Type))
	}
	
	if req.TokenName != nil {
		u.Set("token_name", *req.TokenName)
	}
	
	if req.StartTimestamp != nil {
		u.Set("start_timestamp", fmt.Sprintf("%d", *req.StartTimestamp))
	}
	if req.EndTimestamp != nil {
		u.Set("end_timestamp", fmt.Sprintf("%d", *req.EndTimestamp))
	}
	if req.Channel != nil {
		u.Set("channel", fmt.Sprintf("%d", *req.Channel))
	}

	var r *http.Request
	if r, err = api.createGetRequest(ctx, api.buildRequestApi(PATH_ALL_LOGS), u); err != nil {
		return
	}

	var _resp GetAllLogsResp
	err = api.c.SetHttpRequest(r).SendRequest(&_resp)
	resp = &_resp
	return
}
