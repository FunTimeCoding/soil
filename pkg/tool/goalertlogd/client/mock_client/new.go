package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{}
	result.failures = failure.New(result.GetStatus)
	t.Cleanup(result.Close)

	return result
}
