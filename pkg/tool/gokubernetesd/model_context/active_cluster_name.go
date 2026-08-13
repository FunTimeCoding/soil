package model_context

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) activeClusterName(x context.Context) (string, error) {
	session := server.ClientSessionFromContext(x)

	if session == nil {
		return "", fmt.Errorf("no session")
	}

	v, okay := s.sessions.Load(session.SessionID())

	if !okay {
		return "", fmt.Errorf("no cluster selected - use use_cluster first")
	}

	return v.(string), nil
}
