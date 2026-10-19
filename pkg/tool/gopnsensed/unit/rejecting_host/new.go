package rejecting_host

import "github.com/funtimecoding/soil/pkg/tool/gopnsensed/mock_client"

func New() *Client {
	return &Client{Client: mock_client.New()}
}
