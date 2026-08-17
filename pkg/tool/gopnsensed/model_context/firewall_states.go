package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) firewallStates(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, e := s.opnsense.States(r.GetString(constant.ParameterQuery, ""))

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(convert.States(result))
}
