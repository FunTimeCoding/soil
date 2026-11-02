package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
)

func New() *Client {
	result := &Client{statuses: map[string]*client.Status{}}
	result.failures = failure.New(result.GetStatus)

	return result
}
