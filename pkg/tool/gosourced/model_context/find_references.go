package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) findReferences(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.FindReferences,
) (*mcp.CallToolResult, error) {
	if a.PackagePath == "" {
		return response.Fail("package_path is required")
	}

	if a.Symbol == "" && a.File == "" {
		return response.Fail("pass symbol or file")
	}

	if a.Symbol != "" && a.File != "" {
		return response.Fail("pass symbol or file, not both")
	}

	if a.Receiver != "" && a.Symbol == "" {
		return response.Fail("receiver requires symbol")
	}

	if a.Offset < 0 {
		return response.Fail("offset must not be negative")
	}

	directory, e := s.resolveDirectory(x)

	if e != nil {
		return response.Fail("%s", e)
	}

	if a.File != "" {
		r, references, f := s.service.FileReferences(
			directory,
			a.PackagePath,
			a.File,
		)

		if f != nil {
			return s.captureFail(f, constant.UnexpectedError)
		}

		if references == nil {
			return response.Fail("%s", formatConcerns(r.Entries))
		}

		for _, entry := range references.Symbols {
			entry.Paginate(int(a.Limit), int(a.Offset))
		}

		return response.SuccessAny(references)
	}

	r, references, f := s.service.FindReferences(
		directory,
		a.PackagePath,
		a.Symbol,
		a.Receiver,
	)

	if f != nil {
		return s.captureFail(f, constant.UnexpectedError)
	}

	if references == nil {
		return response.Fail("%s", formatConcerns(r.Entries))
	}

	references.Paginate(int(a.Limit), int(a.Offset))

	return response.SuccessAny(references)
}
