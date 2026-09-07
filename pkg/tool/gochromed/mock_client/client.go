package mock_client

import "github.com/funtimecoding/soil/pkg/chromium/tab"

type Client struct {
	tabs []*tab.Tab
}
