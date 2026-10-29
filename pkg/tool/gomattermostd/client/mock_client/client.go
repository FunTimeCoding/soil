package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"

type Client struct {
	Posts      map[string][]*client.Post
	Reactions  map[string][]string
	Threads    map[string][]*client.Post
	Unresolved map[string][]client.Thread
	Failure    error
	Closed     int
}
