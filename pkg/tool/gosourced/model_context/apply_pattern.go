package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) applyPattern(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ApplyPattern,
) (*mcp.CallToolResult, error) {
	if a.PackagePath == "" {
		return response.Fail("package_path is required")
	}

	if a.Symbol == "" {
		return response.Fail("symbol is required")
	}

	if a.Pattern == "" {
		return response.Fail("pattern is required")
	}

	if a.Replacement == "" {
		return response.Fail("replacement is required")
	}

	directory, e := s.resolveDirectory(x)

	if e != nil {
		return response.Fail("%s", e)
	}

	r, apply, f := s.service.ApplyPattern(
		directory,
		a.PackagePath,
		a.Symbol,
		a.Receiver,
		a.Pattern,
		a.Replacement,
		a.Partial,
		a.DryRun,
	)

	if f != nil {
		return s.captureFail(f, constant.UnexpectedError)
	}

	if apply == nil {
		return response.Fail("%s", formatConcerns(r.Entries))
	}

	return response.SuccessAny(apply)
}
