package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/model_context/tab_response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListTabs(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListTabs,
) (*mcp.CallToolResult, error) {
	tabs, e := s.service.Tabs()

	if e != nil {
		return s.captureDetail(x, e)
	}

	result := []*tab_response.Response{}

	for _, t := range tabs {
		result = append(result, tab_response.New(t.Identifier, t.Locator))
	}

	return response.SuccessAny(result)
}
