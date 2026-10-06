package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) introduceConstructor(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.IntroduceConstructor,
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

	r, c, f := s.service.IntroduceConstructor(
		directory,
		a.PackagePath,
		a.Type,
		a.Parameters,
		a.ParametersFromSites,
		a.AssignRest,
		a.DryRun,
	)

	if f != nil {
		return s.captureFail(f, constant.UnexpectedError)
	}

	if c == nil {
		return response.Fail("%s", formatConcerns(r.Entries))
	}

	return response.SuccessAny(c)
}
