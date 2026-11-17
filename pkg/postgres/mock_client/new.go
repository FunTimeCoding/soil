package mock_client

import (
	"github.com/funtimecoding/soil/pkg/postgres/types/result"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	c := &Client{results: make(map[string]*result.Result)}
	t.Cleanup(c.Close)

	return c
}
