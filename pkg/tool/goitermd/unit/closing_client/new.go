package closing_client

import "github.com/funtimecoding/soil/pkg/tool/goitermd/mock_client"

func New() *Client {
	return &Client{Client: mock_client.New()}
}
