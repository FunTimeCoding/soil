package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) useModule(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.UseModule,
) (*mcp.CallToolResult, error) {
	if a.Module == "" {
		return response.Fail("module is required")
	}

	if _, okay := s.service.Module(a.Module); !okay {
		return response.Fail("unknown module: %s", a.Module)
	}

	session := server.ClientSessionFromContext(x)

	if session == nil {
		return response.Fail("no session")
	}

	s.service.SetActiveModule(session.SessionID(), a.Module)

	return response.Success("active module set to %s", a.Module)
}
