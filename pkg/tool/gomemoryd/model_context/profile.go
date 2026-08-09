package model_context

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/convert"
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
		if errors.Is(e, constant.ErrorReservedScope) {
			return response.Fail("%s", e.Error())
		}

		message := "failed to load profile"

		if errors.Is(e, constant.ErrorAlwaysLoad) {
			message = "failed to load always-tagged memories"
		} else if errors.Is(e, constant.ErrorRelevantSearch) {
			message = "failed to search for relevant memories"
		} else if errors.Is(e, constant.ErrorMemoryList) {
			message = "failed to list memories"
		}

		return s.captureFail(e, message)
	}

	return response.Success(
		notation.MarshalIndent(
			profileResponse{
				Always:      convert.Memories(result.Always),
				Relevant:    convert.SearchResults(result.Relevant),
				Index:       result.Index,
				Impressions: result.Impressions,
				Completions: result.Completions,
			},
		),
	)
}
