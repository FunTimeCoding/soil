package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateChoiceSet(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ChoiceSet,
) (*mcp.CallToolResult, error) {
	if a.Name == "" {
		return response.Fail("name is required")
	}

	choices := a.Choices

	if choices == nil {
		choices = []string{}
	}

	result, e := s.service.CreateChoiceSet(a.Name, choices)

	if e != nil {
		return s.answerMessage(x, e, "choice set not created")
	}

	return response.SuccessAny(convert.ChoiceSet(result))
}
