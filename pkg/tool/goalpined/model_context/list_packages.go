package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/convert"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listPackages(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListPackages,
) (*mcp.CallToolResult, error) {
	listings, e := s.service.Packages(a.Name)

	if e != nil {
		return s.answer(e)
	}

	return response.SuccessAny(convert.Listings(listings))
}
