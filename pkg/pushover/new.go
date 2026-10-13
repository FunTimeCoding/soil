package pushover

import "github.com/funtimecoding/soil/pkg/web/stall_client"

func New(
	user string,
	token string,
) *Client {
	return &Client{user: user, token: token, pool: stall_client.New().Build()}
}
