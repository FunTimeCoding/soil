package mock_client

import (
	"github.com/funtimecoding/soil/pkg/browser/chromium/mock_client"
	"github.com/funtimecoding/soil/pkg/browser/face"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{
		Client: mock_client.New(t),
		pages:  make(map[string]face.ChromiumProtocolClient),
	}
	t.Cleanup(result.Close)

	return result
}
