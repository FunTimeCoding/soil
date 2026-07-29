package server

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/mark3labs/mcp-go/server"
	"net/http"
	"time"
)

func (s *Server) Setup(m *http.ServeMux) {
	h := http.NewServeMux()
	h.Handle(
		constant.LocationModelContext,
		server.NewStreamableHTTPServer(
			s.server,
			server.WithStreamableHTTPLogger(s.Logger()),
			server.WithHeartbeatInterval(15*time.Second),
		),
	)
	sse := server.NewSSEServer(s.server)
	h.Handle(constant.LocationEvent, sse.SSEHandler())
	h.Handle(constant.LocationEventMessage, sse.MessageHandler())
	s.wrapAuthentication(m, h)
}
