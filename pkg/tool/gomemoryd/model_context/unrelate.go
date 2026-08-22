package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) unrelate(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	sourceID := int64(q.GetFloat(constant.SourceIdentifier, 0))
	targetID := int64(q.GetFloat(constant.TargetIdentifier, 0))

	if sourceID == 0 || targetID == 0 {
		return response.Fail("source_id and target_id are required")
	}

	removed, e := s.service.DeleteRelation(sourceID, targetID)

	if e != nil {
		return s.captureFail(e, "remove relation")
	}

	if !removed {
		return response.Fail(
			"no relation from %d to %d - the row is directional; try the reverse order",
			sourceID,
			targetID,
		)
	}

	return response.Success(
		fmt.Sprintf("Unrelated memory %d ↔ %d", sourceID, targetID),
	)
}
