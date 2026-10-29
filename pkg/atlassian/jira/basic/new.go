package basic

import (
	"github.com/funtimecoding/soil/pkg/atlassian/classify_refusal"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/basic"
)

func New(
	base *locator.Locator,
	user string,
	token string,
) *Client {
	return &Client{
		requester: requester.New(base).
			WithAuthorizer(basic.New(user, token)).
			WithHeader(constant.Accept, constant.Object).
			WithRefusal(classify_refusal.Refusal),
		base: base,
	}
}
