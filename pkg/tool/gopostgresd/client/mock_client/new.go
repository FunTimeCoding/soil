package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"

func New() *Client {
	return &Client{Result: &client.QueryResult{}}
}
