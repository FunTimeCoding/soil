package model_context

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/mark3labs/mcp-go/server"
	"sync"
)

type Server struct {
	server            *server.MCPServer
	client            *chromium.Client
	downloadDirectory string
	snapshotCache     map[string]map[string]int64
	mutex             sync.Mutex
	reporter          face.Reporter
}
