package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) restoreBackup(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.RestoreBackup,
) (*mcp.CallToolResult, error) {
	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	database, e := s.backup.Restore(x, instance, a.Backup, a.Target)

	if e != nil {
		return s.captureDetail(x, e)
	}

	return response.SuccessAny(server.RestoreResult{Database: database})
}
