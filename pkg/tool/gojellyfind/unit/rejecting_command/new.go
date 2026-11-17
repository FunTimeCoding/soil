package rejecting_command

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/mock_client"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{Client: mock_client.New(t)}
	t.Cleanup(result.Close)

	return result
}
