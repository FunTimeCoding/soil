package model_context

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) resolveDirectory(x context.Context) (string, error) {
	session := server.ClientSessionFromContext(x)

	if session == nil {
		return "", fmt.Errorf("no session")
	}

	name, okay := s.service.ActiveModule(session.SessionID())

	if !okay {
		return "", fmt.Errorf(
			"no active module - call use_module first",
		)
	}

	m, okay := s.service.Module(name)

	if !okay {
		return "", fmt.Errorf("unknown module: %s", name)
	}

	return m.Directory, nil
}
