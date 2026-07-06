package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) useInstance(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.UseInstance,
) (*mcp.CallToolResult, error) {
	if a.Instance == "" {
		return response.Fail("instance is required")
	}

	if _, okay := s.service.Instance(a.Instance); !okay {
		return response.Fail("unknown instance: %s", a.Instance)
	}

	session := server.ClientSessionFromContext(x)

	if session == nil {
		return response.Fail("no session")
	}

	s.service.SetActiveInstance(session.SessionID(), a.Instance)

	return response.Success("active instance set to %s", a.Instance)
}
