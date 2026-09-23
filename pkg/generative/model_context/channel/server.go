package channel

import (
	"github.com/mark3labs/mcp-go/server"
	"sync"
	"time"
)

type Server struct {
	server    *server.MCPServer
	sink      Sink
	sleep     func(time.Duration)
	nonce     string
	mutex     sync.Mutex
	callsign  string
	ready     chan struct{}
	open      chan struct{}
	readyOnce sync.Once
	openOnce  sync.Once
}
