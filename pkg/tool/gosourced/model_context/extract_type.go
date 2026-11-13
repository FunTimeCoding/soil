package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) extractType(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ExtractType,
) (*mcp.CallToolResult, error) {
	if a.PackagePath == "" {
		return response.Fail("package_path is required")
	}

	if a.Type == "" {
		return response.Fail("type is required")
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

	r, e := s.symbol.ExtractType(
		turnContext(x),
		directory,
		a.PackagePath,
		a.Type,
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
