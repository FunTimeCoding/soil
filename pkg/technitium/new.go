package technitium

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
)

func New(
	host string,
	token string,
) *Client {
	errors.FatalOnEmpty(host, "host")
	errors.FatalOnEmpty(token, "token")

	return &Client{
		base:   join.Empty(locator.New(host).String(), "/api"),
		token:  token,
		client: &http.Client{},
	}
}
