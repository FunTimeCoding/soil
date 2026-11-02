package server

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"time"
)

func (s *Server) Setup(g *guard.Mux) {
	h := server.NewStreamableHTTPServer(
		s.server,
		server.WithStreamableHTTPLogger(s.Logger()),
		server.WithHeartbeatInterval(15*time.Second),
		server.WithSessionIdleTTL(constant.ModelContextSessionIdle),
		server.WithStreamableHTTPProtocolVersions(
			mcp.LegacyProtocolVersions()...,
		),
	)
	sse := server.NewSSEServer(s.server)
	s.register(g, webConstant.ModelContextPath, h)
	s.register(g, webConstant.EventPath, sse.SSEHandler())
	s.register(g, webConstant.EventMessagePath, sse.MessageHandler())

	if s.openAuthentication {
		g.Open(webConstant.ProtectedResource, s.protectedResource)
	}
}
