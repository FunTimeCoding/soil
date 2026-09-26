package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	port int,
	token string,
) *Client {
	c, e := client.NewClientWithResponses(
		locator.New(host).Port(port).Insecure().String(),
		client.WithRequestEditorFn(web.BearerEditor(token)),
	)
	errors.PanicOnError(e)

	return &Client{context: context.Background(), client: c}
}
