package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listTables(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListTables,
) (*mcp.CallToolResult, error) {
	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(e)
	}

	v, e := s.service.ListTables(x, instance, a.Schema)

	if e != nil {
		return s.answerWith(
			e,
			fmt.Sprintf(
				"database error on %s: tables in %s not listed",
				instance,
				a.Schema,
			),
		)
	}

	return response.SuccessAny(v)
}
