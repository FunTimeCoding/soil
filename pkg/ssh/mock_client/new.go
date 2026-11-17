package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{
		files:       make(map[string][]byte),
		directories: make(map[string]bool),
	}
	result.failures = failure.New(result.ReadFile)
	t.Cleanup(result.Close)

	return result
}
