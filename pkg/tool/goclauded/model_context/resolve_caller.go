package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/telemetry/constant"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/caller"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) resolveCaller(
	x context.Context,
	tool string,
) (*caller.Caller, error) {
	session := server.ClientSessionFromContext(x)

	if session == nil {
		return caller.New(), nil
	}

	modelContextSessionIdentifier := session.SessionID()
	name, sessionIdentifier, e := s.service.ResolveModelContextSession(
		modelContextSessionIdentifier,
	)

	if e != nil {
		return nil, e
	}

	s.telemetry.Record(
		record.NewDomain(tool, constant.ModelContext, name, constant.Success),
	)
	c := caller.New()
	c.Callsign = name
	c.SessionIdentifier = sessionIdentifier

	return c, nil
}
