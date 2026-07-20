package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/telemetry/constant"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) resolveCaller(
	x context.Context,
	tool string,
) (*caller, error) {
	session := server.ClientSessionFromContext(x)

	if session == nil {
		return &caller{}, nil
	}

	modelContextSessionIdentifier := session.SessionID()
	name, sessionIdentifier, e := s.service.ResolveModelContextSession(modelContextSessionIdentifier)

	if e != nil {
		return nil, e
	}

	s.telemetry.Record(
		record.NewDomain(
			tool,
			constant.ModelContext,
			name,
			constant.Success,
		),
	)

	return &caller{Callsign: name, SessionIdentifier: sessionIdentifier}, nil
}
