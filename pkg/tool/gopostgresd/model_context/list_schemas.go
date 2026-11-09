package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listSchemas(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListSchemas,
) (*mcp.CallToolResult, error) {
	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	v, e := s.catalog.ListSchemas(x, instance)

	if e != nil {
		return s.answerMessage(
			x,
			e,
			fmt.Sprintf("database error on %s: schemas not listed", instance),
		)
	}

	return response.SuccessAny(v)
}
