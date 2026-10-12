package model_context

import (
	"context"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) activeInstanceName(x context.Context) string {
	session := server.ClientSessionFromContext(x)

	if session == nil {
		return ""
	}

	result, _ := s.service.ActiveInstance(session.SessionID())

	return result
}
