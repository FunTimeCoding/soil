package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{next: 1}
	result.failures = failure.New(result.GetViews)
	t.Cleanup(result.Close)

	return result
}
