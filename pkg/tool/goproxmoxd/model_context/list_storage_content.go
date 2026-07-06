package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	proxResponse "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListStorageContent(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListStorageContent,
) (*mcp.CallToolResult, error) {
	if a.Node == "" {
		return response.Fail("node is required")
	}

	if a.Storage == "" {
		return response.Fail("storage is required")
	}

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	content, e := s.service.ListStorageContent(c, a.Node, a.Storage)

	if e != nil {
		return s.captureDetail(e)
	}

	rows := make([]proxResponse.StorageContent, len(content))

	for i, v := range content {
		rows[i] = proxResponse.StorageContent{
			Volume: v.Volid,
			Format: v.Format,
			Size:   v.Size,
		}
	}

	return response.SuccessAny(rows)
}
