package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"testing"
)

func New(t *testing.T, address string) *Client {
	t.Helper()
	result := &Client{address: address}
	result.failures = failure.New(result.PostTrigger)
	t.Cleanup(result.Close)

	return result
}
