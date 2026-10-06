package technitium

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/technitium/basic"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	token string,
) *Client {
	errors.FatalOnEmpty(host, "host")
	errors.FatalOnEmpty(token, "token")

	return &Client{basic: basic.New(locator.New(host), token)}
}
