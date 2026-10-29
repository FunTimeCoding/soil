package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"

func New() *Client {
	return &Client{HostValue: &client.Host{}}
}
