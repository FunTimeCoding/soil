package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
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

	if e := s.selection.Select(x, a.Module); e != nil {
		return s.captureDetail(e)
	}

	return response.Success("active module set to %s", a.Module)
}
