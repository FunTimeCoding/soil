package basic

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/jellyfin/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/header"
)

func New(
	base *locator.Locator,
	token string,
) *Client {
	return &Client{
		requester: requester.New(base).WithAuthorizer(
			header.New(
				web.Authorization,
				fmt.Sprintf(constant.TokenFormat, token),
			),
		),
	}
}
