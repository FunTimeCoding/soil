package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listBackups(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListBackups,
) (*mcp.CallToolResult, error) {
	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	v, e := s.backup.List(instance)

	if e != nil {
		return s.captureDetail(x, e)
	}

	return response.SuccessAny(convert.Backups(v))
}
