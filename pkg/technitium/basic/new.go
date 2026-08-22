package basic

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
)

func New(
	host string,
	token string,
) *Client {
	return &Client{
		base:   join.Empty(locator.New(host).String(), "/api"),
		token:  token,
		client: &http.Client{},
	}
}
