package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) addHost(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, e := s.opnsense.AddHost(hostRequest(r))

	if e != nil {
		return s.answer(e)
	}

	if r.GetBool(constant.ParameterApply, true) {
		if f := s.opnsense.ReconfigureDnsmasq(); f != nil {
			return s.answer(f)
		}
	}

	return response.Success(identifier)
}
