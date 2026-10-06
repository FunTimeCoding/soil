package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/removal"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) removeParameters(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.RemoveParameters,
) (*mcp.CallToolResult, error) {
	if len(a.Functions) == 0 {
		return response.Fail("functions is required")
	}

	var removals []*removal.Parameter

	for _, f := range a.Functions {
		if f.PackagePath == "" || f.Name == "" {
			return response.Fail("every function needs package_path and name")
		}

		if len(f.Parameters) == 0 {
			return response.Fail("%s lists no parameters", f.Name)
		}

		removals = append(
			removals,
			removal.NewParameter(
				f.PackagePath,
				f.Name,
				f.Receiver,
				f.Parameters,
			),
		)
	}

	directory, e := s.resolveDirectory(x)

	if e != nil {
		return response.Fail("%s", e)
	}

	r, e := s.service.RemoveParameters(directory, removals, a.DryRun)

	if e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	var unfixed []*concern.Concern
	var fixed []*concern.Concern

	for _, c := range r.Entries {
		if c.Fixed || c.Planned {
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
