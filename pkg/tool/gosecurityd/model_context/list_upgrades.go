package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListUpgrades(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListUpgrades,
) (*mcp.CallToolResult, error) {
	result, e := s.system.Upgrades(a.Show, a.Score)

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(convert.UpgradeList(result))
}
