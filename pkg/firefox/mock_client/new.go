package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/firefox/tab_group"
)

func New() *Client {
	result := &Client{groups: map[int]*tab_group.Group{}, nextIdentifier: 1}
	result.failures = failure.New(result.Tabs)

	return result
}
