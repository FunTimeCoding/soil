package client

import (
	"github.com/funtimecoding/soil/pkg/nextcloud/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/header"
)

func New(
	host string,
	port int,
	secure bool,
	token string,
) *Client {
	l := locator.New(host).Port(port)

	if !secure {
		l.Insecure()
	}

	return &Client{
		requester: requester.New(l).WithAuthorizer(
			header.New(constant.TokenHeader, token),
		),
	}
}
