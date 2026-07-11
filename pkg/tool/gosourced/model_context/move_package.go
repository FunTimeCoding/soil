package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/lint/concern"
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
		return response.Fail("%s", e)
	}

	r, e := s.service.MovePackage(
		directory,
		a.PackagePath,
		a.TargetPackagePath,
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
