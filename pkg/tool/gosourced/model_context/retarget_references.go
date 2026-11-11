package model_context

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) retargetReferences(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.RetargetReferences,
) (*mcp.CallToolResult, error) {
	directory, e := s.resolveDirectory(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	var moves []*path_move.Move

	for _, m := range a.Paths {
		moves = append(moves, path_move.New(m.From, m.To))
	}

	r, e := s.reference.RetargetReferences(
		turnContext(x),
		directory,
		moves,
		a.DryRun,
	)

	if e != nil {
		return s.captureFail(x, e, soil.UnexpectedError)
	}

	return answerConcerns(r)
}
