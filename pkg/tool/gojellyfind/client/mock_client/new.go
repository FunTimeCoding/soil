package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
)

func New() *Client {
	result := &Client{
		Episodes:  map[string][]*client.Item{},
		Tracks:    map[string][]*client.Item{},
		Played:    map[string][]client.PlayJSONRequestBody{},
		Commanded: map[string][]client.PlaybackCommandJSONRequestBody{},
		Volumes:   map[string][]int{},
	}
	result.failures = failure.New(result.GetItem)

	return result
}
