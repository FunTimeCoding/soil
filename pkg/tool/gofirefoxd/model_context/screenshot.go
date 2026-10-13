package model_context

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"os"
	"path/filepath"
)

func (s *Server) Screenshot(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Screenshot,
) (*mcp.CallToolResult, error) {
	v, e := s.client.Screenshot(int(a.TabIdentifier))

	if e != nil {
		return s.captureDetail(e)
	}

	b, e := base64.StdEncoding.DecodeString(v.Content)

	if e != nil {
		return s.captureFail(e, "screenshot decode failed")
	}

	path := filepath.Join(
		s.downloadDirectory,
		fmt.Sprintf("firefox_%d.png", v.Identifier),
	)

	if f := os.WriteFile(path, b, 0644); f != nil {
		return s.captureFail(f, "file write failed")
	}

	return response.SuccessAny(
		map[string]any{"path": path, "tab": v.Title, "size": len(b)},
	)
}
