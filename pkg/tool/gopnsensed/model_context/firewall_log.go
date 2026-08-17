package model_context

import (
	"context"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) firewallLog(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, e := s.opnsense.Log(
		int(r.GetFloat(generative.ParameterLimit, constant.DefaultLogLimit)),
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(convert.LogEntries(result))
}
