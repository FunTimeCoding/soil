package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	token string,
) *Client {
	errors.FatalOnEmpty(host, constant.Host)
	errors.FatalOnEmpty(token, "token")

	return &Client{
		basic: basic.New(locator.New(host).Base(constant.Base).Trail(), token),
	}
}
