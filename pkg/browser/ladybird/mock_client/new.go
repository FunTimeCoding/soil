package mock_client

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/browser/types/message"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{
		results:     make(map[string]json.RawMessage),
		refusals:    make(map[string]error),
		navigations: make(map[string][]*message.Reply),
		image:       []byte("png"),
	}
	t.Cleanup(result.Close)

	return result
}
