package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) updateChoiceSet(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, e := r.RequireString(generativeConstant.ParameterName)

	if e != nil {
		return response.Fail("name is required: %v", e)
	}

	result, f := s.client.SetChoiceSetChoices(
		name,
		r.GetStringSlice(constant.Choices, []string{}),
	)

	if not_found.Is(f) {
		return response.Fail("%v", f)
	}

	if f != nil {
		return s.captureFail(f, "choice set not updated")
	}

	return response.SuccessAny(convert.ChoiceSet(result))
}
