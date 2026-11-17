package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/file"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{
		files:  make(map[string]*file.File),
		traces: make(map[int64]string),
	}
	t.Cleanup(result.Close)

	return result
}
