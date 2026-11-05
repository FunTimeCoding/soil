package limit_recorder

import "github.com/funtimecoding/soil/pkg/opnsense/mock_client"

func New() *Client {
	return &Client{Client: mock_client.New()}
}
