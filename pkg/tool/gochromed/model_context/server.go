package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	chromium "github.com/funtimecoding/soil/pkg/tool/gochromed/face"
	"github.com/mark3labs/mcp-go/server"
	"sync"
)

type Server struct {
	server            *server.MCPServer
	client            chromium.ChromiumSource
	downloadDirectory string
	snapshotCache     map[string]map[string]int64
	mutex             sync.Mutex
	reporter          face.Reporter
}
