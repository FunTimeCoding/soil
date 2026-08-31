package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) matchPattern(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.MatchPattern,
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

	directory, e := s.resolveDirectory(x)

	if e != nil {
		return response.Fail("%s", e)
	}

	r, match, f := s.service.MatchPattern(
		directory,
		a.PackagePath,
		a.Symbol,
		a.Receiver,
		a.Pattern,
	)

	if f != nil {
		return s.captureFail(f, constant.UnexpectedError)
	}

	if match == nil {
		return response.Fail("%s", formatConcerns(r.Entries))
	}

	return response.SuccessAny(match)
}
