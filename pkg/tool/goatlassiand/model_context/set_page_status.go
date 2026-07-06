package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) setPageStatus(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, f := r.RequireString(parameter.Identifier)

	if f != nil {
		return response.Fail("identifier is required: %v", f)
	}

	status, g := r.RequireString(constant.PageStatus)

	if g != nil {
		return response.Fail("status is required: %v", g)
	}

	if status != "current" && status != "draft" {
		return response.Fail(
			"status must be 'current' (publish) or 'draft' (unpublish)",
		)
	}

	result, h := s.service.SetPageStatus(identifier, status)

	if h != nil {
		return s.captureFail(h, "page status not updated")
	}

	return response.SuccessAny(result)
}
