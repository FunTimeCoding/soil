package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chromium/tab"
	"sync"
)

type Client struct {
	tabs    []*tab.Tab
	closed  int
	Failure error
	mutex   sync.Mutex
}
