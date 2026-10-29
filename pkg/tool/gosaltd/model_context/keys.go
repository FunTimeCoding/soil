package model_context

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) keys(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Keys()

	if errors.Is(e, constant.ErrorNotConnected) {
		return response.Fail(constant.NotConnected)
	}

	if e != nil {
		return s.captureFail(e, constant.KeysFail)
	}

	return response.SuccessAny(result)
}
