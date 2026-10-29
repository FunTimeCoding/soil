package mock_client

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"

func New() *Client {
	return &Client{documents: map[string]*client.SearchResult{}}
}
