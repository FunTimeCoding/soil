package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateCustomField(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateCustomField,
) (*mcp.CallToolResult, error) {
	if a.Name == "" || a.ChoiceSet == "" {
		return response.Fail("name and choice set are required")
	}

	result, e := s.service.CreateSelectionField(a.Name, a.Label, a.ChoiceSet)

	if e != nil {
		return s.answerMessage(x, e, "custom field not created")
	}

	return response.SuccessAny(convert.CustomField(result))
}
