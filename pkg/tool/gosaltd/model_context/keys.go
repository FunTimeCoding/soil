package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) keys(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c := s.runner.SaltClient()

	if c == nil {
		return response.Fail(constant.NotConnected)
	}

	result, e := c.Keys()

	if e != nil {
		return s.captureFail(e, constant.KeysFailed)
	}

	return response.SuccessAny(result)
}
