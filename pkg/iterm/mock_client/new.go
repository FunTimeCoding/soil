package mock_client

import (
	"github.com/funtimecoding/soil/pkg/iterm/screen"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{
		screens:        map[string]*screen.Screen{},
		keys:           map[string][]string{},
		nextIdentifier: 1,
	}
	t.Cleanup(result.Close)

	return result
}
