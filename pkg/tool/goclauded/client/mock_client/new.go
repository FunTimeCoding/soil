package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/target"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{sessions: []*target.Target{}}
	t.Cleanup(result.Close)

	return result
}
