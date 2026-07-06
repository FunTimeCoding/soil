package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
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
		return response.Fail("%s", e)
	}

	r, e := s.service.AddImport(directory, a.File, a.ImportPath, a.Alias)

	if e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	return response.Success(formatConcerns(r.Entries))
}
