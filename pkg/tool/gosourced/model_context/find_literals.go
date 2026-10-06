package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) findLiterals(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.FindLiterals,
) (*mcp.CallToolResult, error) {
	if a.PackagePath == "" {
		return response.Fail("package_path is required")
	}

	if a.Type == "" {
		return response.Fail("type is required")
	}

	directory, e := s.resolveDirectory(x)

	if e != nil {
		return response.Fail("%s", e)
	}

	r, literals, f := s.service.FindLiterals(directory, a.PackagePath, a.Type)

	if f != nil {
		return s.captureFail(f, constant.UnexpectedError)
	}

	if literals == nil {
		return response.Fail("%s", formatConcerns(r.Entries))
	}

	return response.SuccessAny(literals)
}
