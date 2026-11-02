package mock_client

import "github.com/funtimecoding/soil/pkg/failure"

func New(address string) *Client {
	result := &Client{address: address}
	result.failures = failure.New(result.PostTrigger)

	return result
}
