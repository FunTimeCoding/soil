package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"os"
	"path/filepath"
)

func (s *Server) DownloadFile(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.DownloadFile,
) (*mcp.CallToolResult, error) {
	if a.FileIdentifier == "" {
		return response.Fail("file_id is required")
	}

	i, e := s.service.FileInformation(a.FileIdentifier)

	if e != nil {
		return s.answerWith(e, "file information")
	}

	b, e := s.service.File(a.FileIdentifier)

	if e != nil {
		return s.answerWith(e, "file content")
	}

	path := a.Path

	if path == "" {
		path = filepath.Join(os.TempDir(), i.Name)
	}

	e = os.WriteFile(path, b, 0644)

	if e != nil {
		return s.answerWith(e, "write file")
	}

	return response.SuccessAny(
		map[string]any{
			"file_id":   i.Id,
			"name":      i.Name,
			"mime_type": i.MimeType,
			"size":      i.Size,
			"path":      path,
		},
	)
}
