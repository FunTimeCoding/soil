package model_context

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"io/fs"
	"os"
	"path/filepath"
)

func (s *Server) DownloadFile(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DownloadFile,
) (*mcp.CallToolResult, error) {
	if a.FileIdentifier == "" {
		return response.Fail("file_id is required")
	}

	i, e := s.service.FileInformation(a.FileIdentifier)

	if e != nil {
		return s.answerMessage(x, e, "file information")
	}

	b, e := s.service.File(a.FileIdentifier)

	if e != nil {
		return s.answerMessage(x, e, "file content")
	}

	path := a.Path

	if path == "" {
		path = filepath.Join(os.TempDir(), i.Name)
	}

	e = os.WriteFile(path, b, 0644)

	if errors.Is(e, fs.ErrNotExist) {
		return s.answerMessage(
			x,
			validation.New("directory does not exist: %s", filepath.Dir(path)),
			"write file",
		)
	}

	if e != nil {
		return s.answerMessage(x, e, "write file")
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
