package mock_client

import (
	"github.com/funtimecoding/soil/pkg/docker/hub/tag"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{tags: map[string][]*tag.Tag{}}
	t.Cleanup(result.Close)

	return result
}
