package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chromium/tab"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/face"
)

type Client struct {
	tabs  []*tab.Tab
	pages map[string]face.Page
}
