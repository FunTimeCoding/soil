package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
)

func New() *Client {
	result := &Client{
		Posts:      map[string][]*client.Post{},
		Reactions:  map[string][]string{},
		Threads:    map[string][]*client.Post{},
		Unresolved: map[string][]client.Thread{},
	}
	result.failures = failure.New(result.GetPostReactions)

	return result
}
