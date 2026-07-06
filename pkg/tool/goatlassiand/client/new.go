package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	port int,
	insecure bool,
) *Client {
	l := locator.New(host).Port(port)

	if insecure {
		l.Insecure()
	}

	c, e := client.NewClientWithResponses(l.String())
	errors.PanicOnError(e)

	return &Client{context: context.Background(), client: c}
}
