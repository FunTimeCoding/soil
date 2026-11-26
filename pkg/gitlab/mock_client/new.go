package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/file"
	"gitlab.com/gitlab-org/api/client-go/v3"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{
		files:     make(map[string]*file.File),
		traces:    make(map[int64]string),
		approvals: make(map[[2]int64][]string),
		threads:   make(map[string][]*gitlab.Note),
		resolved:  make(map[string]bool),
	}
	t.Cleanup(result.Close)

	return result
}
