package web

import (
	"crypto/tls"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net"
	"net/http"
)

func newStallTransport(c *tls.Config) *http.Transport {
	return &http.Transport{
		Proxy:             http.ProxyFromEnvironment,
		ForceAttemptHTTP2: true,
		TLSClientConfig:   c,
		DialContext: (&net.Dialer{
			Timeout:   constant.DialTimeout,
			KeepAlive: constant.KeepAlive,
		}).DialContext,
		TLSHandshakeTimeout:   constant.HandshakeTimeout,
		ResponseHeaderTimeout: constant.ResponseHeaderTimeout,
		IdleConnTimeout:       constant.IdleTimeout,
		MaxIdleConns:          constant.MaximumIdle,
	}
}
