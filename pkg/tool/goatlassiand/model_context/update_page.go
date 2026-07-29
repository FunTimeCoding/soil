package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) updatePage(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, f := r.RequireString(constant.ParameterIdentifier)

	if f != nil {
		return response.Fail("identifier is required: %v", f)
	}

	title, g := r.RequireString(constant.ParameterTitle)

	if g != nil {
		return response.Fail("title is required: %v", g)
	}

	body, h := r.RequireString(constant.ParameterBody)

	if h != nil {
		return response.Fail("body is required: %v", h)
	}

	message := r.GetString(constant.ParameterMessage, "")
	result, i := s.service.UpdatePage(identifier, title, body, message)

	if i != nil {
		return s.captureFail(i, "page not updated")
	}

	return response.SuccessAny(result)
}
