package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{statuses: map[string]*client.Status{}}
	result.failures = failure.New(result.GetStatus)
	t.Cleanup(result.Close)

	return result
}
