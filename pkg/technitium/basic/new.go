package basic

import (
	"github.com/funtimecoding/soil/pkg/technitium/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/bearer"
)

func New(
	root *locator.Locator,
	token string,
) *Client {
	base := root.Copy().Base(constant.Base)

	return &Client{
		requester: requester.New(base).WithAuthorizer(bearer.New(token)),
		base:      base.String(),
	}
}
