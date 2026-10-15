package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ansibleInventory(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, e := s.service.AnsibleInventory(
		r.GetString(constant.Field, constant.AnsibleField),
	)

	if e != nil {
		return s.captureFail(e, "inventory not built")
	}

	return response.Success(result.Markup())
}
