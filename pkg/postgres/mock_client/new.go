package mock_client

import "github.com/funtimecoding/soil/pkg/postgres/types/result"

func New() *Client {
	return &Client{results: make(map[string]*result.Result)}
}
