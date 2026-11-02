package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/lint/types/symbol_move"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) retargetReferences(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.RetargetReferences,
) (*mcp.CallToolResult, error) {
	var paths []*path_move.Move

	for _, m := range a.Paths {
		paths = append(paths, path_move.New(m.From, m.To))
	}

	var symbols []*symbol_move.Move

	for _, m := range a.Symbols {
		symbols = append(
			symbols,
			symbol_move.New(m.Directory, m.Name, m.NewDirectory, m.NewName),
		)
	}

	reports, e := s.service.RetargetReferences(
		move_set.New(paths, symbols),
		a.DryRun,
	)

	if e != nil {
		return s.answer(e)
	}

	if len(reports) == 0 {
		return response.Success(constant.RetargetClean)
	}

	return response.Success(join.NewLine(retargetLines(reports)))
}
