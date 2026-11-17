package mock_client

import (
	"github.com/funtimecoding/soil/pkg/github/run"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{runs: make(map[string][]*run.Run)}
	t.Cleanup(result.Close)

	return result
}
