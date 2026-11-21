package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetUpgrade(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.GetUpgrade,
) (*mcp.CallToolResult, error) {
	result, e := s.system.Upgrade(a.Release, a.Source)

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(convert.UpgradeGroup(result))
}
