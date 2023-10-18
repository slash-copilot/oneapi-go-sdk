package oneapigosdk

import (
	"net/http"
	"time"
)

type ClientConfig struct {
	Host        string
	AccessToken string
	Timeout     time.Duration
	Transport   *http.Transport
}
