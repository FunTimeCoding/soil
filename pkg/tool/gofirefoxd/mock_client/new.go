package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gofirefoxd/types/tab_group"

func New() *Client {
	return &Client{groups: map[int]*tab_group.Group{}, nextIdentifier: 1}
}
