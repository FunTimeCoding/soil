package model_context

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) extractToFile(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ExtractToFile,
) (*mcp.CallToolResult, error) {
	if a.File == "" {
		return response.Fail("file is required")
	}

	if a.Symbol == "" {
		return response.Fail("symbol is required")
	}

	directory, e := s.resolveDirectory(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	r, e := s.file.ExtractToFile(
		turnContext(x),
		directory,
		a.File,
		a.Symbol,
		a.DryRun,
	)

	if e != nil {
		return s.captureFail(x, e, soil.UnexpectedError)
	}

	return answerConcerns(r)
}
