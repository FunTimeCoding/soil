package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) moveSymbol(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.MoveSymbol,
) (*mcp.CallToolResult, error) {
	if a.PackagePath == "" {
		return response.Fail("package_path is required")
	}

	if a.Symbol == "" {
		return response.Fail("symbol is required")
	}

	if a.TargetPackagePath == "" {
		return response.Fail("target_package_path is required")
	}

	if a.PackagePath == a.TargetPackagePath {
		return response.Fail(
			"package_path and target_package_path are the same",
		)
	}

	directory, e := s.resolveDirectory(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	r, e := s.symbol.MoveSymbol(
		turnContext(x),
		directory,
		a.PackagePath,
		a.Symbol,
		a.TargetPackagePath,
		a.TargetFile,
		a.Create,
		a.DryRun,
	)

	if e != nil {
		return s.captureFail(x, e)
	}

	return answerConcerns(r)
}
