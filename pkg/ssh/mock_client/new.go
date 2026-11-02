package mock_client

import "github.com/funtimecoding/soil/pkg/failure"

func New() *Client {
	result := &Client{
		files:       make(map[string][]byte),
		directories: make(map[string]bool),
	}
	result.failures = failure.New(result.ReadFile)

	return result
}
