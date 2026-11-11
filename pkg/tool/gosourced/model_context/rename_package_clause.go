package model_context

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) renamePackageClause(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.RenamePackageClause,
) (*mcp.CallToolResult, error) {
	if a.PackagePath == "" {
		return response.Fail("package_path is required")
	}

	if a.NewName == "" {
		return response.Fail("new_name is required")
	}

	directory, e := s.resolveDirectory(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	r, e := s.packages.RenamePackageClause(
		turnContext(x),
		directory,
		a.PackagePath,
		a.NewName,
		a.DryRun,
	)

	if e != nil {
		return s.captureFail(x, e, soil.UnexpectedError)
	}

	return answerConcerns(r)
}
