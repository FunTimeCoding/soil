package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) describeTable(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DescribeTable,
) (*mcp.CallToolResult, error) {
	if a.Table == "" {
		return response.Fail("table is required")
	}

	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(e)
	}

	v, e := s.service.DescribeTable(x, instance, a.Schema, a.Table)

	if e != nil {
		return s.answerWith(
			e,
			fmt.Sprintf(
				"database error on %s: table %s.%s not described",
				instance,
				a.Schema,
				a.Table,
			),
		)
	}

	return response.SuccessAny(v)
}
