package model_context

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/receiver_rename"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) renameReceivers(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.RenameReceivers,
) (*mcp.CallToolResult, error) {
	if a.PackagePath == "" {
		return response.Fail("package_path is required")
	}

	var renames []*receiver_rename.Rename

	for _, n := range a.Receivers {
		if n.Type == "" || n.Name == "" {
			return response.Fail("every receiver needs type and name")
		}

		renames = append(renames, receiver_rename.New(n.Type, n.Name))
	}

	directory, e := s.resolveDirectory(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	r, e := s.symbol.RenameReceivers(
		turnContext(x),
		directory,
		a.PackagePath,
		renames,
		a.DryRun,
	)

	if e != nil {
		return s.captureFail(x, e, soil.UnexpectedError)
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
