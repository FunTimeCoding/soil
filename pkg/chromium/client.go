package chromium

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"sync"
)

type Client struct {
	host            string
	port            int
	requester       *requester.Requester
	allocator       context.Context
	allocatorCancel context.CancelFunc
	context         context.Context
	cancel          context.CancelFunc
	targets         map[string]context.Context
	mutex           sync.Mutex
}
