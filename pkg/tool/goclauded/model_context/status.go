package model_context

import (
	"context"
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) status(
	x context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	if _, e := s.resolveCaller(x, constant.Status); e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	findings, e := s.service.Findings()

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	if len(findings) == 0 {
		return response.Success("Nothing inconsistent.")
	}

	var lines []string

	for _, i := range findings {
		line := fmt.Sprintf("[%s] %s", i.Kind, i.Detail)

		if i.Subject != "" {
			line = fmt.Sprintf("[%s] %s: %s", i.Kind, i.Subject, i.Detail)
		}

		lines = append(lines, line)
	}

	return response.Success(join.NewLine(lines))
}
