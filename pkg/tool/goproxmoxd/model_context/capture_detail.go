package model_context

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/luthermonson/go-proxmox"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
	if errors.Is(e, proxmox.ErrNotAuthorized) ||
		errors.Is(e, proxmox.ErrTimeout) {
		return s.captureFail(e, e.Error())
	}

	return response.Answer(s.reporter, e)
}
