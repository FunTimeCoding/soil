package model_context

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) movePackage(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.MovePackage,
) (*mcp.CallToolResult, error) {
	if a.PackagePath == "" {
		return response.Fail("package_path is required")
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

	r, e := s.packages.MovePackage(
		turnContext(x),
		directory,
		a.PackagePath,
		a.TargetPackagePath,
		a.DryRun,
	)

	if e != nil {
		return s.captureFail(x, e, soil.UnexpectedError)
	}

	return answerConcerns(r)
}
