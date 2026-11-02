package client

import "github.com/funtimecoding/soil/pkg/web/locator"

func New(
	l *locator.Locator,
	token string,
) *Client {
	return newOnToken(l, func() string { return token })
}
