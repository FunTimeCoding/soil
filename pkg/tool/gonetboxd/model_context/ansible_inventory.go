package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) AnsibleInventory(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.AnsibleInventory,
) (*mcp.CallToolResult, error) {
	field := a.Field

	if field == "" {
		field = constant.AnsibleField
	}

	result, e := s.service.AnsibleInventory(field, a.Site)

	if e != nil {
		return s.answerMessage(x, e, "inventory not built")
	}

	return response.Success(result.Markup())
}
