package rejecting_command

import "github.com/funtimecoding/soil/pkg/tool/gojellyfind/mock_client"

func New() *Client {
	return &Client{Client: mock_client.New()}
}
