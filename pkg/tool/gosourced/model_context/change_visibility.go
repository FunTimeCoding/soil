package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) changeVisibility(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ChangeVisibility,
) (*mcp.CallToolResult, error) {
	if a.Symbol == "" {
		return response.Fail("symbol is required")
	}

	if a.PackagePath == "" {
		return response.Fail("package_path is required")
	}

	directory, e := s.resolveDirectory(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	r, e := s.symbol.ChangeVisibility(
		turnContext(x),
		directory,
		a.Symbol,
		a.PackagePath,
		a.Receiver,
		a.DryRun,
	)

	if e != nil {
		return s.captureFail(x, e)
	}

	return answerConcerns(r)
}
