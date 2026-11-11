package model_context

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) moveMethods(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.MoveMethods,
) (*mcp.CallToolResult, error) {
	if a.PackagePath == "" || a.Type == "" {
		return response.Fail("package_path and type are required")
	}

	if a.TargetPackagePath == "" || a.TargetType == "" {
		return response.Fail("target_package_path and target_type are required")
	}

	if len(a.Methods) == 0 {
		return response.Fail("methods is required")
	}

	directory, e := s.resolveDirectory(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	r, e := s.symbol.MoveMethods(
		turnContext(x),
		directory,
		a.PackagePath,
		a.Type,
		a.Methods,
		a.TargetPackagePath,
		a.TargetType,
		a.TargetFile,
		a.DryRun,
	)

	if e != nil {
		return s.captureFail(x, e, soil.UnexpectedError)
	}

	return answerConcerns(r)
}
