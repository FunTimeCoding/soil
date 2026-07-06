package mock_client

import "github.com/funtimecoding/soil/pkg/firefox/tab"

type Client struct {
	tabs    []*tab.Tab
	groups  map[int]*group
	nextID  int
	groupID int
}
