package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/web/constant"
	transport "github.com/go-openapi/runtime/client"
	"github.com/prometheus/alertmanager/api/v2/client"
)

func New(
	host string,
	secure bool,
	user string,
	password string,
	p *prometheus.Client,
) *Client {
	errors.FatalOnEmpty(host, "host")
	scheme := constant.Insecure

	if secure {
		scheme = constant.Secure
	}

	c := client.NewHTTPClient(nil)
	t := transport.New(
		host,
		client.DefaultBasePath,
		[]string{scheme},
	)

	if user != "" && password != "" {
		t.DefaultAuthentication = transport.BasicAuth(user, password)
	}

	c.SetTransport(t)

	return &Client{client: c, host: host, prometheus: p}
}
