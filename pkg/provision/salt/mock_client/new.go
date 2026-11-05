package mock_client

import "github.com/funtimecoding/soil/pkg/failure"

func New() *Client {
	result := &Client{}
	result.failures = failure.New(result.Keys)

	return result
}
