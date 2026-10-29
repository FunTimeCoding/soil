package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createPage(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	space, e := r.RequireString(constant.SpaceIdentifier)

	if e != nil {
		return response.Fail("space_identifier is required: %v", e)
	}

	parent, f := r.RequireString(constant.ParentIdentifier)

	if f != nil {
		return response.Fail("parent_identifier is required: %v", f)
	}

	draft := r.GetBool(constant.Draft, false)
	title := r.GetString(generativeConstant.ParameterTitle, "")

	if !draft && title == "" {
		return response.Fail("title is required for published pages")
	}

	body, g := r.RequireString(generativeConstant.ParameterBody)

	if g != nil {
		return response.Fail("body is required: %v", g)
	}

	result, h := s.service.CreatePage(space, parent, title, body, draft)

	if h != nil {
		return s.answer(h)
	}

	return response.SuccessAny(result)
}
