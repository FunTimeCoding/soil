package mock_client

import "github.com/funtimecoding/soil/pkg/github/run"

func New() *Client {
	return &Client{runs: make(map[string][]*run.Run)}
}
