package mock_client

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/mock_client/page_entry"
	"github.com/funtimecoding/soil/pkg/failure"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{pages: map[string]*page_entry.Entry{}, nextIdentifier: 1}
	result.failures = failure.New(result.Page)
	t.Cleanup(result.Close)

	return result
}
