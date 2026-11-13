package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) renameModule(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.RenameModule,
) (*mcp.CallToolResult, error) {
	if a.ModulePath == "" {
		return response.Fail("module_path is required")
	}

	if a.NewModulePath == "" {
		return response.Fail("new_module_path is required")
	}

	directory, e := s.resolveDirectory(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	r, e := s.module.RenameModule(
		turnContext(x),
		directory,
		a.ModulePath,
		a.NewModulePath,
		a.Force,
		a.DryRun,
	)

	if e != nil {
		return s.captureFail(x, e)
	}

	return answerConcerns(r)
}
