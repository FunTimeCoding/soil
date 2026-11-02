package mock_client

import "github.com/funtimecoding/soil/pkg/docker/hub/tag"

func New() *Client {
	return &Client{tags: map[string][]*tag.Tag{}}
}
