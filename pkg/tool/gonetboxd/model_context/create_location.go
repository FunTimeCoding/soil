package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateLocation(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateLocation,
) (*mcp.CallToolResult, error) {
	if a.Name == "" || a.Site == "" {
		return response.Fail("name and site are required")
	}

	result, e := s.service.CreateLocation(a.Name, a.Site)

	if e != nil {
		return s.answerMessage(x, e, "location not created")
	}

	return response.SuccessAny(convert.Location(result))
}
