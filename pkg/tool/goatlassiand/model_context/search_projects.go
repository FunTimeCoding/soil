package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) searchProjects(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	limit, f := r.RequireFloat(generativeConstant.ParameterLimit)

	if f != nil {
		return response.Fail("limit is required: %v", f)
	}

	matches, h := s.service.SearchProjects(
		r.GetString(generativeConstant.ParameterQuery, ""),
		int(limit),
		r.GetBool(constant.IncludeDescriptions, false),
	)

	if h != nil {
		return s.answer(h)
	}

	var result []*server.JiraProject

	for _, m := range matches {
		if m.Detail != nil {
			result = append(result, convert.JiraProject(m.Detail))

			continue
		}

		o := &server.JiraProject{Key: m.Key, Name: m.Name}

		if m.Type != "" {
			o.Type = &m.Type
		}

		result = append(result, o)
	}

	return response.SuccessAny(result)
}
