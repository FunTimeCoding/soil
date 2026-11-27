package mock_client

import (
	"github.com/funtimecoding/soil/pkg/browser/firefox/tab_group"
	"github.com/funtimecoding/soil/pkg/failure"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{groups: map[int]*tab_group.Group{}, nextIdentifier: 1}
	result.failures = failure.New(result.Tabs)
	t.Cleanup(result.Close)

	return result
}
