package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

func (s *Server) History(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.History,
) (*mcp.CallToolResult, error) {
	t, e := s.resolveTab(a.TabID, a.Title, a.URL)

	if e != nil {
		return response.Fail(e.Error())
	}

	h, e := s.client.History(t.Identifier)

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
			entry.URL,
		)
	}

	return response.Success(b.String())
}
