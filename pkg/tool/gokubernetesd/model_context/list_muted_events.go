package model_context

import (
	"context"
	mark "github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListMutedEvents(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListMutedEvents,
) (*mcp.CallToolResult, error) {
	rules, e := s.service.MuteRules()

	if e != nil {
		return s.captureFail(e, "list mute rules")
	}

	result := []response.MuteRuleEntry{}

	for _, rule := range rules {
		result = append(
			result,
			response.MuteRuleEntry{
				Identifier: rule.Identifier,
				Reason:     rule.Reason,
				Message:    rule.Message,
				Cluster:    rule.Cluster,
			},
		)
	}

	return mark.SuccessAny(result)
}
