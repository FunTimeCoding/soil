package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/event"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/label"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/memory"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/roster"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/search"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/usage"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server    *server.MCPServer
	roster    *roster.Service
	label     *label.Service
	message   *message.Service
	event     *event.Service
	memory    *memory.Service
	usage     *usage.Service
	search    *search.Service
	reporter  face.Reporter
	logger    *logger.Logger
	telemetry face.Recorder
}
