package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getPage(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, f := r.RequireString(parameter.Identifier)

	if f != nil {
		return response.Fail("identifier is required: %v", f)
	}

	var result *page.Page
	var g error

	if r.GetBool(constant.Draft, false) {
		result, g = s.confluence.DraftPage(identifier)
	} else {
		result, g = s.confluence.Page(identifier)
	}

	if g != nil {
		return s.captureFail(g, "page not found")
	}

	return response.SuccessAny(result)
}
