package basic

import (
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/basic"
)

func New(
	fileRoot *locator.Locator,
	user string,
	password string,
) *Client {
	return &Client{
		requester: requester.New(fileRoot).WithAuthorizer(
			basic.New(user, password),
		),
	}
}
