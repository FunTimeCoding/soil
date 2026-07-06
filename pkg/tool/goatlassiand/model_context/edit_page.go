package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

func (s *Server) editPage(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, f := r.RequireString(parameter.Identifier)

	if f != nil {
		return response.Fail("identifier is required: %v", f)
	}

	oldText, g := r.RequireString(constant.OldText)

	if g != nil {
		return response.Fail("old_text is required: %v", g)
	}

	newText, h := r.RequireString(constant.NewText)

	if h != nil {
		return response.Fail("new_text is required: %v", h)
	}

	title := r.GetString(parameter.Title, "")
	message := r.GetString(parameter.Message, "")
	draft := r.GetBool(constant.Draft, false)
	result, i := s.service.EditPage(
		identifier,
		oldText,
		newText,
		title,
		message,
		draft,
	)

	if i != nil {
		if strings.Contains(i.Error(), constant.OldText) {
			return response.Fail("%s", i)
		}

		return s.captureFail(i, "page not updated")
	}

	return response.SuccessAny(result)
}
