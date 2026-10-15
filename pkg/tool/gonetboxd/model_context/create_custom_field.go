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

func (s *Server) createCustomField(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, e := r.RequireString(generativeConstant.ParameterName)

	if e != nil {
		return response.Fail("name is required: %v", e)
	}

	set, f := r.RequireString(constant.ChoiceSet)

	if f != nil {
		return response.Fail("choice_set is required: %v", f)
	}

	result, g := s.client.CreateSelectionField(
		name,
		r.GetString(constant.Label, ""),
		set,
	)

	if not_found.Is(g) {
		return response.Fail("%v", g)
	}

	if g != nil {
		return s.captureFail(g, "custom field not created")
	}

	return response.SuccessAny(convert.CustomField(result))
}
