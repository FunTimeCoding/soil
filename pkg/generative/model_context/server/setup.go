package server

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/mark3labs/mcp-go/server"
	"time"
)

func (s *Server) Setup(g *guard.Mux) {
	h := server.NewStreamableHTTPServer(
		s.server,
		server.WithStreamableHTTPLogger(s.Logger()),
		server.WithHeartbeatInterval(15*time.Second),
	)
	sse := server.NewSSEServer(s.server)
	s.register(g, constant.ModelContextPath, h)
	s.register(g, constant.EventPath, sse.SSEHandler())
	s.register(g, constant.EventMessagePath, sse.MessageHandler())

	if s.openAuthentication {
		g.Open(constant.ProtectedResource, s.protectedResource)
	}
}
