package model_context

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) renameSymbol(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.RenameSymbol,
) (*mcp.CallToolResult, error) {
	if a.PackagePath == "" {
		return response.Fail("package_path is required")
	}

	if a.OldName == "" {
		return response.Fail("old_name is required")
	}

	if a.NewName == "" {
		return response.Fail("new_name is required")
	}

	if a.OldName == a.NewName {
		return response.Fail("old_name and new_name are the same")
	}

	directory, e := s.resolveDirectory(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	r, e := s.symbol.Rename(
		turnContext(x),
		directory,
		a.PackagePath,
		a.OldName,
		a.NewName,
		a.Receiver,
		a.DryRun,
	)

	if e != nil {
		return s.captureFail(x, e, soil.UnexpectedError)
	}

	return answerConcerns(r)
}
