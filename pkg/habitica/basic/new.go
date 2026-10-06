package basic

import (
	"github.com/funtimecoding/soil/pkg/habitica/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/header"
)

func New(
	base *locator.Locator,
	userIdentifier string,
	token string,
) *Client {
	return &Client{
		requester: requester.New(base).
			WithHeader(constant.UserHeader, userIdentifier).
			WithAuthorizer(header.New(constant.TokenHeader, token)),
	}
}
