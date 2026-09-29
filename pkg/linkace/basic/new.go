package basic

import "github.com/funtimecoding/soil/pkg/web/locator"

func New(
	host string,
	token string,
) *Client {
	return &Client{base: locator.New(host).Base("api/v2"), token: token}
}
