package basic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/bearer"
)

func New(
	base *locator.Locator,
	token string,
) *Client {
	return &Client{
		requester: requester.New(base).
			WithAuthorizer(bearer.New(token)).
			WithHeader(constant.Accept, constant.Object),
	}
}
