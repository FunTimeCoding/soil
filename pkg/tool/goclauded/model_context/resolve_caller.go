package model_context

import (
	"context"
	telemetry "github.com/funtimecoding/soil/pkg/telemetry/constant"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
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

	s.logger.Structured(
		"model_context_tool_call",
		"tool",
		tool,
		"model_context_session_identifier",
		modelContextSessionIdentifier,
		constant.SessionName,
		name,
	)
	s.telemetry.Record(
		record.NewDomain(
			tool,
			telemetry.ModelContext,
			name,
			telemetry.Success,
		),
	)

	return &caller{Callsign: name, SessionIdentifier: sessionIdentifier}, nil
}
