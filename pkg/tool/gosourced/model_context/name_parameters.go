package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/naming"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) nameParameters(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.NameParameters,
) (*mcp.CallToolResult, error) {
	if a.PackagePath == "" || a.Type == "" || a.Name == "" {
		return response.Fail("package_path, type and name are required")
	}

	var functions []*naming.Function

	for _, f := range a.Functions {
		if f.Name == "" {
			return response.Fail("every function needs a name")
		}

		functions = append(functions, naming.NewFunction(f.Name, f.Receiver))
	}

	directory, e := s.resolveDirectory(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	r, e := s.symbol.NameParameters(
		turnContext(x),
		directory,
		naming.NewRequest(a.PackagePath, functions, a.Type, a.Name, a.Imports),
		a.DryRun,
	)

	if e != nil {
		return s.captureFail(x, e)
	}

	return answerConcerns(r)
}
