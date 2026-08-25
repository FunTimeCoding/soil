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
	sourceIdentifier := int64(q.GetFloat(constant.SourceIdentifier, 0))
	targetIdentifier := int64(q.GetFloat(constant.TargetIdentifier, 0))

	if sourceIdentifier == 0 || targetIdentifier == 0 {
		return response.Fail("source_id and target_id are required")
	}

	removed, e := s.service.DeleteRelation(sourceIdentifier, targetIdentifier)

	if e != nil {
		return s.captureFail(e, "remove relation")
	}

	if !removed {
		return response.Fail(
			"no relation from %d to %d - the row is directional; try the reverse order",
			sourceIdentifier,
			targetIdentifier,
		)
	}

	return response.Success(
		fmt.Sprintf(
			"Unrelated memory %d ↔ %d",
			sourceIdentifier,
			targetIdentifier,
		),
	)
}
