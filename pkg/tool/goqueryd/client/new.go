package client

import (
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func New(
	l *locator.Locator,
	token string,
) *Client {
	return wrap(l, token, stall_client.New().Build())
}
