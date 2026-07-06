package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
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
	title := r.GetString(parameter.Title, "")

	if !draft && title == "" {
		return response.Fail("title is required for published pages")
	}

	body, g := r.RequireString(parameter.Body)

	if g != nil {
		return response.Fail("body is required: %v", g)
	}

	if draft {
		result, h := s.confluence.CreateDraftPage(space, parent, title, body)

		if h != nil {
			return s.captureFail(h, "draft page not created")
		}

		return response.SuccessAny(result)
	}

	result, h := s.confluence.CreatePage(space, parent, title, body)

	if h != nil {
		return s.captureFail(h, "page not created")
	}

	return response.SuccessAny(result)
}
