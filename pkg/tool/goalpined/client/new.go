package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	token string,
) *Client {
	base := locator.New(host).String()
	c, e := client.NewClientWithResponses(
		base,
		client.WithRequestEditorFn(web.BearerEditor(token)),
	)
	errors.PanicOnError(e)

	return &Client{
		context: context.Background(),
		client:  c,
		base:    base,
		token:   token,
	}
}
