package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

func (s *Server) History(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.History,
) (*mcp.CallToolResult, error) {
	h, e := s.service.History(selector.New(a.TabIdentifier, a.Title, a.Locator))

	if e != nil {
		return s.captureDetail(e)
	}

	var b strings.Builder

	for i, entry := range h.Entries {
		marker := "  "

		if int64(i) == h.CurrentIndex {
			marker = "→ "
		}

		writer.Print(
			&b,
			"%s[%d] %s\n     %s\n",
			marker,
			i,
			entry.Title,
			entry.Locator,
		)
	}

	return response.Success(b.String())
}
