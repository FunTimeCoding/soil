package stall_client

import (
	"crypto/tls"
	"github.com/funtimecoding/soil/pkg/web/basic_transport"
	"github.com/funtimecoding/soil/pkg/web/bearer_transport"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"github.com/funtimecoding/soil/pkg/web/verbose_transport"
	"net"
	"net/http"
)

func (b *Builder) Build() *pool.Pool {
	t := &http.Transport{
		Proxy:             http.ProxyFromEnvironment,
		ForceAttemptHTTP2: true,
		DialContext: (&net.Dialer{
			Timeout:   constant.DialTimeout,
			KeepAlive: constant.KeepAlive,
		}).DialContext,
		TLSHandshakeTimeout:   constant.HandshakeTimeout,
		ResponseHeaderTimeout: constant.ResponseHeaderTimeout,
		IdleConnTimeout:       constant.IdleTimeout,
		MaxIdleConns:          constant.MaximumIdle,
		HTTP2: &http.HTTP2Config{
			SendPingTimeout: b.sendPing,
			PingTimeout:     b.ping,
		},
	}

	if b.insecure {
		t.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	if b.long {
		t.ResponseHeaderTimeout = constant.LongResponseHeaderTimeout
	}

	var r http.RoundTripper = t

	if b.basic {
		r = basic_transport.New(b.user, b.password, r)
	}

	if b.bearer {
		r = bearer_transport.New(b.token, r)
	}

	if b.verbose {
		r = verbose_transport.New(r)
	}

	return pool.New(t, r, b.timeout)
}
