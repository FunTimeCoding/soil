package model_context

import (
	"context"
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) searchConversations(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.SearchConversations,
) (*mcp.CallToolResult, error) {
	if _, e := s.resolveCaller(x, constant.SearchConversations); e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	found, e := s.service.SearchConversations(a.Query, a.Kinds, int(a.Limit))

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	var lines []string

	if indexed, total := s.service.SearchProgress(); indexed < total {
		lines = append(
			lines,
			fmt.Sprintf(
				"indexing %d/%d conversations - results may be incomplete",
				indexed,
				total,
			),
		)
	}

	if len(found) == 0 {
		return response.Success(
			join.NewLine(append(lines, "No conversation holds every term.")),
		)
	}

	for _, c := range found {
		name := c.Name

		if name == "" {
			name = constant.UnnamedSession
		}

		lines = append(
			lines,
			fmt.Sprintf(
				"%s (%s) - %d hits, latest %s",
				name,
				c.Session,
				c.Count,
				c.Latest,
			),
		)

		for _, h := range c.Hits {
			lines = append(
				lines,
				fmt.Sprintf(
					"  [%s %s %s %s] %s",
					h.At,
					h.Role,
					h.Kind,
					h.Identifier,
					h.Snippet,
				),
			)
		}
	}

	return response.Success(join.NewLine(lines))
}
