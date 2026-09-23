package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) setCruise(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SetCruise,
) (*mcp.CallToolResult, error) {
	mode := constant.Cruise(a.Mode)

	switch mode {
	case constant.CruiseOff, constant.CruiseOn, constant.CruisePaced:
	default:
		return response.Fail("mode must be off, on or paced - got %s", a.Mode)
	}

	pace := time.Duration(a.Pace) * time.Minute

	if mode == constant.CruisePaced && pace <= 0 {
		return response.Fail("paced mode needs a pace in minutes")
	}

	s.service.SetCruise(a.Session, mode, pace)

	return response.Success("cruise %s", mode)
}
