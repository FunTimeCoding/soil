package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) create(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, e := q.RequireString(constant.MemoryName)

	if e != nil {
		return response.Fail("name is required")
	}

	content, f := q.RequireString(constant.Content)

	if f != nil {
		return response.Fail("content is required")
	}

	description, g := q.RequireString(constant.Description)

	if g != nil {
		return response.Fail("description is required")
	}

	var parentIdentifier *int64

	if value := q.GetFloat(constant.ParentIdentifier, 0); value > 0 {
		identifier := int64(value)
		parentIdentifier = &identifier
	}

	o := save_option.New()
	o.Name = name
	o.Content = content
	o.Description = description
	o.Type = q.GetString(constant.Type, "")
	o.Source = q.GetString(constant.Source, "")
	o.ParentIdentifier = parentIdentifier
	stripped := false

	if raw := q.GetString(constant.Tags, ""); raw != "" {
		o.Tags, stripped = splitTags(raw)
	}

	m, h := s.service.CreateMemory(o)

	if h != nil {
		return s.captureDetail(h)
	}

	created := fmt.Sprintf("Created memory %d tags: %v", m.Identifier, m.Tags)

	if !stripped {
		return response.Success(created)
	}

	return response.Success(
		join.NewLine(
			[]string{
				join.Space(constant.TagStripNotice, join.CommaSpace(o.Tags)),
				created,
			},
		),
	)
}
