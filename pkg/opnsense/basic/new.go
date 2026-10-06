package basic

import (
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/basic"
)

func New(
	base *locator.Locator,
	key string,
	secret string,
	untrusted bool,
) *Client {
	result := requester.New(base).
		WithAuthorizer(basic.New(key, secret)).
		WithHeader(constant.Accept, constant.Object)

	if untrusted {
		result.WithClient(web.InsecureStallClient())
	}

	return &Client{requester: result}
}
