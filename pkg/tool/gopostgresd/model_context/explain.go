package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) explain(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Explain,
) (*mcp.CallToolResult, error) {
	if a.SQL == "" {
		return response.Fail("sql is required")
	}

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	prefix := "EXPLAIN"

	if a.Analyze {
		prefix = "EXPLAIN ANALYZE"
	}

	v, e := s.service.Query(x, instance, fmt.Sprintf("%s %s", prefix, a.SQL))

	if e != nil {
		return s.captureFail(
			e,
			fmt.Sprintf(
				"database error on %s: explain not executed",
				instance,
			),
		)
	}

	return response.SuccessAny(v)
}
