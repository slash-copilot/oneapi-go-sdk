package oneapigosdk

import (
	"context"
	"net/http"
)

var (
	searchUsers = "/api/user/search"
)

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
