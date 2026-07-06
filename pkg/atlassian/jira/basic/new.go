package basic

import "github.com/funtimecoding/soil/pkg/web/locator"

func New(
	host string,
	user string,
	token string,
) *Client {
	return &Client{
		user:  user,
		token: token,
		host:  host,
		base:  locator.New(host),
	}
}
