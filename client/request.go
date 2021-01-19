package client

import (
	"github.com/Archer1A/quickClient/handler"
	"time"
)

const (
	DefaultRetry     = 0
	DefaultTimeout   = 30 * time.Second
)

type HttpConfig struct {
	Retries     int
	IgnoreTLS   bool
	Timeout     time.Duration
	HttpHandler *handler.HttpHandler
}

func NewDefaultHttpConfig() *HttpConfig {
	return &HttpConfig{
		Retries:     DefaultRetry,
		IgnoreTLS:   true,
		Timeout:     DefaultTimeout,
		HttpHandler: handler.NewDefaultHttpHandler(),
	}
}

func (hc *HttpConfig) WithRetries(retries int) *HttpConfig {
	hc.Retries = retries
	return hc
}
func (hc *HttpConfig) WithTimeout(timeout time.Duration) *HttpConfig {
	hc.Timeout = timeout
	return hc
}

func (hc *HttpConfig) WithIgnoreTLS(ignore bool) *HttpConfig {
	hc.IgnoreTLS = ignore
	return hc
}

func (hc *HttpConfig) WithHttpHandler(handler *handler.HttpHandler) *HttpConfig {
	hc.HttpHandler = handler
	return hc
}
