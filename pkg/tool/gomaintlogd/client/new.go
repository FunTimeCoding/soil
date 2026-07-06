package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(host string) *Client {
	c, e := client.NewClientWithResponses(locator.New(host).String())
	errors.PanicOnError(e)

	return &Client{context: context.Background(), client: c}
}
