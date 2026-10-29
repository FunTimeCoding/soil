package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) equipGear(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(constant.Key)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	result, e := s.service.Equip(key)

	if e != nil {
		return s.answer(e)
	}

	return response.SuccessAny(convert.Gear(result))
}
