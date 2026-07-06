package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Wake(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Wake,
) (*mcp.CallToolResult, error) {
	t, e := s.resolveTab(a.TabID, a.Title, a.URL)

	if e != nil {
		return response.Fail(e.Error())
	}

	e = s.client.Wake(t.Identifier)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.Success("woke tab: %s", t.Title)
}
