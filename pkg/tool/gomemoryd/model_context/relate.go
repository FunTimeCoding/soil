package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) relate(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	sourceID := int64(q.GetFloat(constant.SourceIdentifier, 0))
	targetID := int64(q.GetFloat(constant.TargetIdentifier, 0))

	if sourceID == 0 || targetID == 0 {
		return response.Fail("source_id and target_id are required")
	}

	relationType := q.GetString(constant.Type, "")
	e := s.service.CreateRelation(sourceID, targetID, relationType)

	if validation.Is(e) {
		return response.Fail("%s", e.Error())
	}

	if e != nil {
		return s.captureFail(e, "create relation")
	}

	if relationType != "" {
		return response.Success(
			fmt.Sprintf(
				"Related memory %d ↔ %d (%s)",
				sourceID,
				targetID,
				relationType,
			),
		)
	}

	return response.Success(
		fmt.Sprintf("Related memory %d ↔ %d", sourceID, targetID),
	)
}
