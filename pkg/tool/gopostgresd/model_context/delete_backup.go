package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deleteBackup(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DeleteBackup,
) (*mcp.CallToolResult, error) {
	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	if f := s.backup.Delete(instance, a.Backup); f != nil {
		return s.captureDetail(x, f)
	}

	return response.Success(
		fmt.Sprintf(constant.BackupDeleted, a.Backup, instance),
	)
}
