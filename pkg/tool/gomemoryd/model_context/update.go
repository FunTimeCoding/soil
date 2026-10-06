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

	existing, f := s.service.GetMemory(int64(identifier))

	if f != nil {
		return s.captureDetail(f)
	}

	o := save_option.New()
	o.Name = textOr(q, constant.MemoryName, existing.Name)
	o.Content = textOr(q, constant.Content, existing.Content)
	o.Description = textOr(q, constant.Description, existing.Description)
	o.Source = q.GetString(constant.Source, "")
	o.Base = baseChange(q)
	m, rewritten, h := s.service.UpdateMemory(int64(identifier), o)

	if h != nil {
		return s.captureDetail(h)
	}

	return response.Success(
		s.withReferences(
			withRewritten(
				fmt.Sprintf("Updated memory %d", m.Identifier),
				rewritten,
			),
			m,
		),
	)
}
