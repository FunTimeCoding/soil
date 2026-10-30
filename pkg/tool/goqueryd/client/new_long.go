package client

import (
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func NewLong(
	l *locator.Locator,
	token string,
) *Client {
	return wrap(l, token, stall_client.New().Long().Build())
}
