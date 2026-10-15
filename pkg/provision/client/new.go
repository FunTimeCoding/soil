package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func New(
	host string,
	port int,
	token string,
) *Client {
	p := stall_client.New().Build()
	c, e := client.NewClientWithResponses(
		locator.New(host).Port(port).Insecure().String(),
		client.WithHTTPClient(p.Client()),
		client.WithRequestEditorFn(web.BearerEditor(token)),
	)
	errors.PanicOnError(e)

	return &Client{context: context.Background(), client: c, pool: p}
}
