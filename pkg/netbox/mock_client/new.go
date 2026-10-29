package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/netbox/network"
)

func New() *Client {
	result := &Client{interfaces: make(map[int32][]*network.Interface)}
	result.failures = failure.New(result.Racks)

	return result
}
