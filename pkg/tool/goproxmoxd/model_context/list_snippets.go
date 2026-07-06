package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	proxResponse "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListSnippets(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListSnippets,
) (*mcp.CallToolResult, error) {
	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	c, e := s.service.SSHClient(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	entries := c.ListDirectory(snippetDirectory)
	rows := make([]proxResponse.Snippet, 0, len(entries))

	for _, v := range entries {
		if v.IsDir() {
			continue
		}

		rows = append(
			rows,
			proxResponse.Snippet{
				Name: v.Name(),
				Size: uint64(v.Size()),
			},
		)
	}

	return response.SuccessAny(rows)
}
