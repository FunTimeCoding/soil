package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getGroup(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier := int64(q.GetFloat(constant.MemoryIdentifier, 0))

	if identifier == 0 {
		return response.Fail("memory_id is required")
	}

	parent, children, e := s.service.GetMemoryGroup(identifier)

	if e != nil {
		return s.captureFail(e, "failed to load memory group")
	}

	return response.Success(
		notation.MarshalIndent(
			map[string]any{
				"parent":   parent,
				"children": children,
			},
		),
	)
}
