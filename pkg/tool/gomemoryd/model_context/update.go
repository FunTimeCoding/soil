package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) update(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, e := q.RequireFloat(constant.MemoryIdentifier)

	if e != nil {
		return response.Fail("memory_id is required")
	}

	name, f := q.RequireString(constant.MemoryName)

	if f != nil {
		return response.Fail("name is required")
	}

	content, g := q.RequireString(constant.Content)

	if g != nil {
		return response.Fail("content is required")
	}

	description, h := q.RequireString(constant.Description)

	if h != nil {
		return response.Fail("description is required")
	}

	m, i := s.service.UpdateMemory(
		int64(identifier),
		name,
		content,
		description,
		q.GetString(constant.Source, ""),
	)

	if i != nil {
		return s.captureDetail(i)
	}

	return response.Success(fmt.Sprintf("Updated memory %d", m.Identifier))
}
