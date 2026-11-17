package mock_client

import (
	"github.com/funtimecoding/soil/pkg/nextcloud/usage"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{usage: usage.New(0, 0)}
	t.Cleanup(result.Close)

	return result
}
