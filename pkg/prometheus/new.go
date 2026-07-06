package prometheus

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/client"
	"github.com/funtimecoding/soil/pkg/prometheus/round_tripper"
	"github.com/funtimecoding/soil/pkg/web"
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

	return &Client{
		context: context.Background(),
		client: client.New(
			web.Link(host, port, secure),
			round_tripper.New(user, password),
		),
		host:               host,
		alternateGraphHost: alternateGraphHost,
	}
}
