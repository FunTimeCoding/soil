package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) minions(
	x context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Minions()

	if unavailable.Is(e) {
		return response.Fail("%v", e)
	}

	if e != nil {
		return s.captureFail(x, e, constant.MinionsFail)
	}

	return response.SuccessAny(result)
}
