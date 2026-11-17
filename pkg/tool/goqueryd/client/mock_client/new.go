package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{documents: map[string]*client.SearchResult{}}
	t.Cleanup(result.Close)

	return result
}
