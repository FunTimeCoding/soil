package mock_client

import (
	"github.com/funtimecoding/soil/pkg/browser/chromium/tab"
	"sync"
)

type Client struct {
	tabs    []*tab.Tab
	closed  int
	Failure error
	mutex   sync.Mutex
}
