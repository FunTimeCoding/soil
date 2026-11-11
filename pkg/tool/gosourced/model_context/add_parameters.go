package model_context

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/addition"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) addParameters(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.AddParameters,
) (*mcp.CallToolResult, error) {
	if len(a.Functions) == 0 {
		return response.Fail("functions is required")
	}

	var additions []*addition.Parameter

	for _, f := range a.Functions {
		if f.PackagePath == "" || f.Name == "" {
			return response.Fail("every function needs package_path and name")
		}

		if f.Parameter == "" || f.Type == "" || f.Argument == "" {
			return response.Fail(
				"%s needs parameter, type and argument",
				f.Name,
			)
		}

		additions = append(
			additions,
			addition.NewParameter(
				f.PackagePath,
				f.Name,
				f.Receiver,
				f.Parameter,
				f.Type,
				f.Position,
				f.Argument,
			),
		)
	}

	directory, e := s.resolveDirectory(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	r, e := s.symbol.AddParameters(
		turnContext(x),
		directory,
		additions,
		a.Imports,
		a.DryRun,
	)

	if e != nil {
		return s.captureFail(x, e, soil.UnexpectedError)
	}

	return answerConcerns(r)
}
