package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol_rename"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) renameSymbols(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.RenameSymbols,
) (*mcp.CallToolResult, error) {
	if a.PackagePath == "" {
		return response.Fail("package_path is required")
	}

	var renames []*symbol_rename.Rename

	for _, n := range a.Renames {
		if n.OldName == "" || n.NewName == "" {
			return response.Fail("every rename needs old_name and new_name")
		}

		if n.OldName == n.NewName {
			return response.Fail(
				"%s: old_name and new_name are the same",
				n.OldName,
			)
		}

		renames = append(
			renames,
			symbol_rename.New(n.OldName, n.NewName, n.Receiver),
		)
	}

	directory, e := s.resolveDirectory(x)

	if e != nil {
		return s.captureDetail(e)
	}

	r, e := s.service.RenameSymbols(
		turnContext(x),
		directory,
		a.PackagePath,
		renames,
		a.DryRun,
	)

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
