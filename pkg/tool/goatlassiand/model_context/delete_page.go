package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deletePage(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, f := r.RequireString(parameter.Identifier)

	if f != nil {
		return response.Fail("identifier is required: %v", f)
	}

	draft := r.GetBool(constant.Draft, false)

	if draft {
		if g := s.confluence.DeleteDraft(identifier); g != nil {
			return s.captureFail(g, "draft page not deleted")
		}
	} else {
		if g := s.confluence.Delete(identifier); g != nil {
			return s.captureFail(g, "page not deleted")
		}
	}

	return response.Success("page deleted")
}
