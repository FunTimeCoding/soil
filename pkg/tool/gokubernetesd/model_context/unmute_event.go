package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) UnmuteEvent(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.UnmuteEvent,
) (*mcp.CallToolResult, error) {
	if a.Identifier == 0 {
		return response.Fail("identifier is required")
	}

	e := s.service.UnmuteEvent(uint(a.Identifier))

	if e != nil {
		return s.captureFail(e, "delete mute rule")
	}

	return response.Success("unmuted rule %d", a.Identifier)
}
