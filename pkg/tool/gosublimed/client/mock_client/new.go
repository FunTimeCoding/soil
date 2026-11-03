package mock_client

import "github.com/funtimecoding/soil/pkg/failure"

func New() *Client {
	result := &Client{next: 1}
	result.failures = failure.New(result.GetViews)

	return result
}
