package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) profile(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, _, e := s.service.Profile(
		q.GetString(constant.Topic, ""),
		q.GetString(constant.Scope, ""),
		false,
	)

	if e != nil {
		if validation.Is(e) {
			return response.Fail("%s", e.Error())
		}

		return s.captureFail(e, "load profile")
	}

	return response.Success(result.Text)
}
