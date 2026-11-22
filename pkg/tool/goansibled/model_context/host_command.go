package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) hostCommand(
	x context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	command, e := r.RequireString(constant.HostCommandParameter)

	if e != nil {
		return response.Fail("command is required: %v", e)
	}

	host, f := r.RequireString(constant.HostCommandHost)

	if f != nil {
		return response.Fail("host is required: %v", f)
	}

	result, g := s.service.HostCommand(
		command,
		host,
		provisionConstant.RunnerTriggerManual,
	)

	if g != nil {
		return s.answer(x, g)
	}

	return response.SuccessAny(result)
}
