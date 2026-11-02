package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/face"
)

func New() *Client {
	result := &Client{pages: make(map[string]face.Page)}
	result.failures = failure.New(result.Tabs)

	return result
}
