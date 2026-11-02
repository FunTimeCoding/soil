package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listIndexes(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListIndexes,
) (*mcp.CallToolResult, error) {
	if a.Table == "" {
		return response.Fail("table is required")
	}

	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(e)
	}

	v, e := s.service.ListIndexes(x, instance, a.Schema, a.Table)

	if e != nil {
		return s.answerWith(
			e,
			fmt.Sprintf(
				"database error on %s: indexes for %s.%s not listed",
				instance,
				a.Schema,
				a.Table,
			),
		)
	}

	return response.SuccessAny(v)
}
