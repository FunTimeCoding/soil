package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) delete(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier := r.GetFloat(constant.Identifier, 0)

	if identifier == 0 {
		return response.Fail("id is required")
	}

	if e := s.service.Delete(uint(identifier)); e != nil {
		return s.answer(e)
	}

	return response.Success("Entry %d deleted successfully", int(identifier))
}
