package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) moveSymbols(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.MoveSymbols,
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

	if len(a.Symbols) == 0 && a.File == "" {
		return response.Fail("pass symbols or file")
	}

	if len(a.Symbols) > 0 && a.File != "" {
		return response.Fail("pass symbols or file, not both")
	}

	directory, e := s.resolveDirectory(x)

	if e != nil {
		return response.Fail("%s", e)
	}

	r, e := s.service.MoveSymbols(
		directory,
		a.PackagePath,
		a.Symbols,
		a.File,
		a.TargetPackagePath,
		a.TargetFile,
		a.Create,
		a.QualifyBackReferences,
	)

	if e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	var unfixed []*concern.Concern
	var fixed []*concern.Concern

	for _, c := range r.Entries {
		if c.Fixed {
			fixed = append(fixed, c)
		} else {
			unfixed = append(unfixed, c)
		}
	}

	if len(unfixed) > 0 {
		return response.Fail("%s", formatConcerns(unfixed))
	}

	return response.Success(formatConcerns(fixed))
}
