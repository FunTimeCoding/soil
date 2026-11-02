package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	transport "github.com/go-openapi/runtime/client"
	"github.com/prometheus/alertmanager/api/v2/client"
)

func New(
	host string,
	port int,
	secure bool,
	user string,
	password string,
	o ...Option,
) *Client {
	errors.FatalOnEmpty(host, "host")
	scheme := constant.Insecure

	if secure {
		scheme = constant.Secure
	}

	connection := locator.NewEndpoint(host, port, secure)
	c := client.NewHTTPClient(nil)
	b := stall_client.New().Build()
	t := transport.New(
		web.HostPortFromLocator(connection.String()),
		client.DefaultBasePath,
		[]string{scheme},
	)
	t.Transport = b.Client().Transport

	if user != "" && password != "" {
		t.DefaultAuthentication = transport.BasicAuth(user, password)
	}

	c.SetTransport(t)
	result := &Client{client: c, link: connection, pool: b}

	for _, f := range o {
		f(result)
	}

	return result
}
