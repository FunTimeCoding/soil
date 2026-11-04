package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) retargetReferences(
	x context.Context,
	_ mcp.CallToolRequest,
	a server.RetargetRequest,
) (*mcp.CallToolResult, error) {
	reports, e := s.service.RetargetReferences(convert.MoveSet(&a), a.DryRun)

	if e != nil {
		return s.answer(x, e)
	}

	if len(reports) == 0 {
		return response.Success(constant.RetargetClean)
	}

	return response.Success(join.NewLine(retargetLines(reports)))
}
