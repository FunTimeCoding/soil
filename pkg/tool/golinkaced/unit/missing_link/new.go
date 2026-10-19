package missing_link

import "github.com/funtimecoding/soil/pkg/tool/golinkaced/mock_client"

func New() *Client {
	return &Client{Client: mock_client.New()}
}
