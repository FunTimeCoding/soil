package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{
		Episodes:  map[string][]*client.Item{},
		Tracks:    map[string][]*client.Item{},
		Played:    map[string][]client.PlayJSONRequestBody{},
		Commanded: map[string][]client.PlaybackCommandJSONRequestBody{},
		Volumes:   map[string][]int{},
	}
	result.failures = failure.New(result.GetItem)
	t.Cleanup(result.Close)

	return result
}
