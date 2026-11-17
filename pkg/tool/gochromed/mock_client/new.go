package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chromium/face"
	"github.com/funtimecoding/soil/pkg/chromium/mock_client"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{
		Client: mock_client.New(t),
		pages:  make(map[string]face.ProtocolClient),
	}
	t.Cleanup(result.Close)

	return result
}
