package habitica

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/habitica/basic"
	"github.com/funtimecoding/soil/pkg/habitica/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	userIdentifier string,
	token string,
) *Client {
	errors.FatalOnEmpty(host, "host")
	errors.FatalOnEmpty(userIdentifier, "user")
	errors.FatalOnEmpty(token, "token")

	return &Client{
		basic: basic.New(
			locator.New(host).Base(constant.Base),
			userIdentifier,
			token,
		),
	}
}
