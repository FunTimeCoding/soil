package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
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

	if guard, g := s.guardDocumentSourced(int64(identifier)); guard != nil {
		return guard, g
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

	o := save_option.New()
	o.Name = name
	o.Content = content
	o.Description = description
	o.Source = q.GetString(constant.Source, "")
	m, i := s.service.UpdateMemory(int64(identifier), o)

	if i != nil {
		return s.captureDetail(i)
	}

	return response.Success(fmt.Sprintf("Updated memory %d", m.Identifier))
}
