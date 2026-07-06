package model_context

import (
	"context"
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

func (s *Server) roster(
	x context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	_, e := s.resolveCaller(x, constant.Roster)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	sessions, e := s.service.ListSessions()

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	if len(sessions) == 0 {
		return response.Success("No active sessions.")
	}

	var lines []string

	for _, session := range sessions {
		line := session.CallsignValue()

		if session.Topic != "" {
			line = fmt.Sprintf("%s - %s", line, session.Topic)
		}

		var details []string

		if session.TurnCount > 0 {
			details = append(
				details,
				fmt.Sprintf("%d turns", session.TurnCount),
			)
		}

		if session.Alias != nil {
			details = append(
				details,
				fmt.Sprintf("alias: %s", *session.Alias),
			)
		}

		if session.Description != "" {
			details = append(details, session.Description)
		}

		if session.FirstMessage != "" {
			details = append(
				details,
				fmt.Sprintf("%q", session.FirstMessage),
			)
		}

		if len(details) > 0 {
			line = fmt.Sprintf(
				"%s\n  %s",
				line,
				strings.Join(details, " · "),
			)
		}

		labels, f := s.service.LabelsBySession(session.Identifier)

		if f != nil {
			return s.captureFail(f, library.UnexpectedError)
		}

		if len(labels) > 0 {
			var pips []string

			for _, l := range labels {
				pips = append(
					pips,
					fmt.Sprintf("(%s:%s)", l.Key, l.Value),
				)
			}

			line = fmt.Sprintf("%s\n  %s", line, join.Space(pips...))
		}

		l, g := s.service.LatestPulse(session.Identifier)

		if g != nil {
			return s.captureFail(g, library.UnexpectedError)
		}

		if l != nil {
			line = fmt.Sprintf(
				"%s\n  pulse: %s",
				line,
				l.Body,
			)
		}

		lines = append(lines, line)
	}

	return response.Success(join.NewLine(lines))
}
