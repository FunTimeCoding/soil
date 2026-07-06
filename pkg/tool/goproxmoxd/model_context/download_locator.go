package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) DownloadLocator(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DownloadLocator,
) (*mcp.CallToolResult, error) {
	if a.Node == "" {
		return response.Fail("node is required")
	}

	if a.Storage == "" {
		return response.Fail("storage is required")
	}

	if a.Content == "" {
		return response.Fail("content is required")
	}

	if a.Filename == "" {
		return response.Fail("filename is required")
	}

	if a.Locator == "" {
		return response.Fail("locator is required")
	}

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	e = s.service.DownloadLocator(
		c,
		a.Node,
		a.Storage,
		a.Content,
		a.Filename,
		a.Locator,
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.Success("downloaded")
}
