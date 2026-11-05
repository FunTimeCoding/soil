package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chromium/face"
	"github.com/funtimecoding/soil/pkg/chromium/mock_client"
)

func New() *Client {
	return &Client{
		Client: mock_client.New(),
		pages:  make(map[string]face.ProtocolClient),
	}
}
