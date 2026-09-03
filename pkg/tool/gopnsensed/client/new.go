package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	port int,
	insecure bool,
	token string,
) *Client {
	l := locator.New(host).Port(port)

	if insecure {
		l.Insecure()
	}

	c, e := client.NewClientWithResponses(
		l.String(),
		client.WithRequestEditorFn(web.BearerEditor(token)),
	)
	errors.PanicOnError(e)

	return &Client{context: context.Background(), client: c}
}
