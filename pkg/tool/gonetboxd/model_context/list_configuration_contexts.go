package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listConfigurationContexts(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, e := s.service.ConfigurationContexts()

	if e != nil {
		return s.captureFail(e, "configuration contexts not listed")
	}

	return response.SuccessAny(convert.ConfigurationContexts(result))
}
