package prometheus

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/client"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
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
	b := stall_client.New()

	if user != "" || password != "" {
		b.BasicAuth(user, password)
	}

	p := b.Build()
	connection := locator.NewEndpoint(host, port, secure)
	result := &Client{
		context: context.Background(),
		client:  client.New(connection.String(), p.Client().Transport),
		pool:    p,
		link:    connection.Copy(),
	}

	for _, f := range o {
		f(result)
	}

	return result
}
