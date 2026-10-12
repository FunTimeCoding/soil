package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/network"

func New() *Client {
	return &Client{interfaces: make(map[int32][]*network.Interface)}
}
