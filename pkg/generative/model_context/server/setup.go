package server

import (
	"github.com/funtimecoding/soil/pkg/web/location"
	"github.com/mark3labs/mcp-go/server"
	"net/http"
	"time"
)

func (s *Server) Setup(m *http.ServeMux) {
	h := http.NewServeMux()
	h.Handle(
		location.ModelContext,
		server.NewStreamableHTTPServer(
			s.server,
			server.WithStreamableHTTPLogger(s.Logger()),
			server.WithHeartbeatInterval(15*time.Second),
		),
	)
	sse := server.NewSSEServer(s.server)
	h.Handle(location.Event, sse.SSEHandler())
	h.Handle(location.EventMessage, sse.MessageHandler())
	s.wrapAuthentication(m, h)
}
