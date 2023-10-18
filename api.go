package oneapigosdk

import (
	"context"
	"errors"
	"net/http"
)

var (
	token = "/api/token"
)

type AddTokenResp struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type Token struct {
	Id             int    `json:"id"`
	UserId         int    `json:"user_id"`
	Key            string `json:"key" gorm:"type:char(48);uniqueIndex"`
	Status         int    `json:"status" gorm:"default:1"`
	Name           string `json:"name" gorm:"index" `
	CreatedTime    int64  `json:"created_time" gorm:"bigint"`
	AccessedTime   int64  `json:"accessed_time" gorm:"bigint"`
	ExpiredTime    int64  `json:"expired_time" gorm:"bigint;default:-1"` // -1 means never expired
	RemainQuota    int    `json:"remain_quota" gorm:"default:0"`
	UnlimitedQuota bool   `json:"unlimited_quota" gorm:"default:false"`
	UsedQuota      int    `json:"used_quota" gorm:"default:0"` // used quota
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

func (api *Api) AddToken(ctx context.Context, req *Token) (resp *AddTokenResp, err error) {
	if req == nil {
		err = errors.New("AddToken.Token Illegal")
		return
	}

	var r *http.Request
	if r, err = api.createPostRequest(ctx, api.buildRequestApi(token), req); err != nil {
		return
	}

	var _resp AddTokenResp

	err = api.c.SetHttpRequest(r).SendRequest(&_resp)
	resp = &_resp
	return
}
