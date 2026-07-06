package model_context

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) profile(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, _, e := s.service.Profile(
		q.GetString(constant.Topic, ""),
		false,
	)

	if e != nil {
		message := "failed to load profile"

		if errors.Is(e, service.ErrorAlwaysLoad) {
			message = "failed to load always-tagged memories"
		} else if errors.Is(e, service.ErrorRelevantSearch) {
			message = "failed to search for relevant memories"
		} else if errors.Is(e, service.ErrorMemoryList) {
			message = "failed to list memories"
		}

		return s.captureFail(e, message)
	}

	return response.Success(notation.MarshalIndent(result))
}
