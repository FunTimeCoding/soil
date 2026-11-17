package mock_client

import (
	"github.com/funtimecoding/soil/pkg/openbao/status"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{
		status:  status.New(true, false, "2.7.1"),
		secrets: map[string]map[string]string{},
	}
	t.Cleanup(result.Close)

	return result
}
