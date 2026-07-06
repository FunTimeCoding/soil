package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"os"
	"path/filepath"
)

func (s *Server) ReadBody(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ReadBody,
) (*mcp.CallToolResult, error) {
	t, e := s.resolveTab(a.TabID, a.Title, a.URL)

	if e != nil {
		return response.Fail(e.Error())
	}

	body := s.client.Body(t.Identifier)
	path := filepath.Join(
		s.downloadDirectory,
		fmt.Sprintf("chrome_%s.html", t.Identifier),
	)
	e = os.WriteFile(path, []byte(body), 0644)

	if e != nil {
		return response.Fail("file write failed: %s", e.Error())
	}

	return response.SuccessAny(
		map[string]any{
			"path": path,
			"tab":  t.Title,
			"size": len(body),
		},
	)
}
