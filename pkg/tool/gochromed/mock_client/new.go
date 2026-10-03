package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gochromed/face"

func New() *Client {
	return &Client{pages: make(map[string]face.Page)}
}
