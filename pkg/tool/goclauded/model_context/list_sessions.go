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

func (s *Server) listSessions(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	_, e := s.resolveCaller(x, constant.ListSessions)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	limit := int(q.GetFloat(constant.Limit, 20))
	offset := int(q.GetFloat(constant.Offset, 0))
	sessions, e := s.service.EnrichedSessions(limit, offset)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	if len(sessions) == 0 {
		return response.Success("No sessions found.")
	}

	var lines []string

	for _, e := range sessions {
		name := e.Slug

		if e.Alias != "" {
			name = e.Alias
		}

		if name == "" {
			name = e.Identifier[:8]
		}

		line := fmt.Sprintf(
			"%s  %-7s  %s  %d lines",
			e.Identifier[:8],
			e.Name,
			name,
			e.Lines,
		)

		if e.TurnCount > 0 {
			line = fmt.Sprintf("%s  %d turns", line, e.TurnCount)
		}

		if e.Active {
			line = fmt.Sprintf("%s  (active)", line)
		}

		if e.Description != "" {
			line = fmt.Sprintf("%s\n  %s", line, e.Description)
		}

		lines = append(lines, line)
	}

	return response.Success(join.NewLine(lines))
}
