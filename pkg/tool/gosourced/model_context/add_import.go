package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) addImport(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.AddImport,
) (*mcp.CallToolResult, error) {
	if a.File == "" {
		return response.Fail("file is required")
	}

	if a.ImportPath == "" {
		return response.Fail("import_path is required")
	}

	directory, e := s.resolveDirectory(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	r, e := s.file.AddImport(
		turnContext(x),
		directory,
		a.File,
		a.ImportPath,
		a.Alias,
		a.DryRun,
	)

	if e != nil {
		return s.captureFail(x, e)
	}

	return response.Success(formatConcerns(r.Entries))
}
