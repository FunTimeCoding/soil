package habitica

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/habitica/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
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
		base:           join.Empty(locator.New(host).String(), constant.Base),
		userIdentifier: userIdentifier,
		token:          token,
		client:         &http.Client{},
	}
}
