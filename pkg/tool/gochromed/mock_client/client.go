package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chromium/tab"
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/face"
	"sync"
)

type Client struct {
	tabs     []*tab.Tab
	pages    map[string]face.Page
	failures *failure.Failure
	mutex    sync.Mutex
}
