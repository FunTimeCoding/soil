package n8n

import "github.com/funtimecoding/soil/pkg/web/stall_client"

func New(
	host string,
	token string,
) *Client {
	return &Client{host: host, token: token, pool: stall_client.New().Build()}
}
