package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) activeClusterName(x context.Context) (string, error) {
	session := server.ClientSessionFromContext(x)

	if session == nil {
		return "", fmt.Errorf("no session")
	}

	v, okay := s.sessions.Load(session.SessionID())

	if !okay {
		return "", not_selected.Format(
			"no cluster selected - use use_cluster first",
		)
	}

	return v.(string), nil
}
