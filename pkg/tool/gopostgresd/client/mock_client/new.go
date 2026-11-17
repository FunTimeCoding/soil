package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{Result: &client.QueryResult{}}
	t.Cleanup(result.Close)

	return result
}
