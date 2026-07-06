package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) setParent(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier := int64(q.GetFloat(constant.MemoryIdentifier, 0))

	if identifier == 0 {
		return response.Fail("memory_id is required")
	}

	var parentIdentifier *int64
	parentID := int64(q.GetFloat(constant.ParentIdentifier, 0))

	if parentID > 0 {
		parentIdentifier = &parentID
	}

	if e := s.service.SetParent(identifier, parentIdentifier); e != nil {
		return s.captureDetail(e)
	}

	if parentIdentifier == nil {
		return response.Success(
			fmt.Sprintf("Unset parent for memory %d", identifier),
		)
	}

	return response.Success(
		fmt.Sprintf(
			"Set parent of memory %d to %d",
			identifier,
			*parentIdentifier,
		),
	)
}
