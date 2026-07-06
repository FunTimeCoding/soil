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

	e := s.habitica.Equip(key)

	if e != nil {
		return s.captureDetail(e)
	}

	result, g := s.habitica.UserGear()

	if g != nil {
		return s.captureDetail(g)
	}

	return response.SuccessAny(convert.Gear(result))
}
