package basic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/basic"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
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
		result.WithClient(stall_client.New().Insecure().Build().Client())
	}

	return &Client{requester: result}
}
