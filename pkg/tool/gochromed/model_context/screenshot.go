package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"os"
	"path/filepath"
)

func (s *Server) Screenshot(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Screenshot,
) (*mcp.CallToolResult, error) {
	t, e := s.resolveTab(a.TabID, a.Title, a.URL)

	if e != nil {
		return response.Fail(e.Error())
	}

	x := s.client.AcquireTarget(t.Identifier)
	b, e := withTimeout(
		constant.TargetTimeout,
		func() ([]byte, error) {
			return s.client.Screenshot(x)
		},
	)

	if e != nil {
		return s.captureDetail(e)
	}

	path := filepath.Join(
		s.downloadDirectory,
		fmt.Sprintf("chrome_%s.png", t.Identifier),
	)
	e = os.WriteFile(path, b, 0644)

	if e != nil {
		return s.captureFail(e, "file write failed")
	}

	return response.SuccessAny(
		map[string]any{
			"path": path,
			"tab":  t.Title,
			"size": len(b),
		},
	)
}
