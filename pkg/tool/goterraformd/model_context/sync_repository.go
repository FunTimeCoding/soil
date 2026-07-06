package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/provision/model_context"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) syncRepository(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	return model_context.SyncResponse(s.runner.Sync())
}
