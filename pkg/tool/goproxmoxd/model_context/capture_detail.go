package model_context

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/luthermonson/go-proxmox"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) captureDetail(x context.Context, e error) (*mcp.CallToolResult, error) {
	if errors.Is(e, proxmox.ErrNotAuthorized) ||
		errors.Is(e, proxmox.ErrTimeout) {
		return s.captureFail(x, e, e.Error())
	}

	return response.Answer(x, s.reporter, e)
}
