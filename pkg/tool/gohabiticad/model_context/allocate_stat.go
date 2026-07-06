package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) allocateStat(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	stat, f := r.RequireString(constant.Stat)

	if f != nil {
		return response.Fail("stat is required: %v", f)
	}

	result, g := s.habitica.Allocate(stat)

	if g != nil {
		return s.captureDetail(g)
	}

	return response.SuccessAny(convert.Stats(result))
}
