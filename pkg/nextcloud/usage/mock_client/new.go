package mock_client

import "github.com/funtimecoding/soil/pkg/nextcloud/usage"

func New() *Client {
	return &Client{usage: usage.New(0, 0)}
}
