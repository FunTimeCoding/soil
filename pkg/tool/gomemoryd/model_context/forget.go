package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) forget(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier := int64(q.GetFloat(constant.MemoryIdentifier, 0))

	if identifier == 0 {
		return response.Fail("memory_id is required")
	}

	e := s.service.ForgetMemory(identifier, q.GetString(constant.Source, ""))

	if e != nil {
		return s.captureDetail(e)
	}

	return response.Success("Forgot memory %d", identifier)
}
