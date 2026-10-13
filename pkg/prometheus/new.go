package prometheus

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func New(
	host string,
	port int,
	secure bool,
	user string,
	password string,
	alternateGraphHost string,
	o ...Option,
) *Client {
	errors.FatalOnEmpty(host, "host")
	b := stall_client.New()

	if user != "" || password != "" {
		b.BasicAuth(user, password)
	}

	p := b.Build()

	return &Client{
		context: context.Background(),
		client: client.New(
			web.Link(host, port, secure),
			p.Client().Transport,
		),
		pool:               p,
		host:               host,
		alternateGraphHost: alternateGraphHost,
	}
}
