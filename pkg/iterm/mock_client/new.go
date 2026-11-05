package mock_client

import "github.com/funtimecoding/soil/pkg/iterm/screen"

func New() *Client {
	return &Client{
		screens:        map[string]*screen.Screen{},
		keys:           map[string][]string{},
		nextIdentifier: 1,
	}
}
