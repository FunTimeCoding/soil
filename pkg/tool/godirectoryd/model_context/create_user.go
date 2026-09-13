package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createUser(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateUser,
) (*mcp.CallToolResult, error) {
	result, e := s.service.CreateUser(
		a.Account,
		a.Name,
		a.Surname,
		a.Mail,
		a.Password,
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(result)
}
