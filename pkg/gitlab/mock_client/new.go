package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/file"

func New() *Client {
	return &Client{files: make(map[string]*file.File)}
}
