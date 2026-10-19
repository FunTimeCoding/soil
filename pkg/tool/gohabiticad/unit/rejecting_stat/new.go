package rejecting_stat

import "github.com/funtimecoding/soil/pkg/tool/gohabiticad/mock_client"

func New() *Client {
	return &Client{Client: mock_client.New()}
}
