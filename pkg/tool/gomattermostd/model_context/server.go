package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	mattermostFace "github.com/funtimecoding/soil/pkg/tool/gomattermostd/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server       *server.MCPServer
	client       mattermostFace.MattermostSource
	monitor      *monitor.Monitor
	store    *store.Store
	indexer  mattermostFace.Indexer
	reporter face.Reporter
}
