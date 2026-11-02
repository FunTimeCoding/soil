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

	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(e)
	}

	v, e := s.service.Explain(x, instance, a.SQL, a.Analyze)

	if e != nil {
		return s.answerWith(
			e,
			fmt.Sprintf("database error on %s: explain not executed", instance),
		)
	}

	return response.SuccessAny(v)
}
