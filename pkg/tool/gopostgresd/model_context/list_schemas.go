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
	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	v, e := s.service.ListSchemas(x, instance)

	if e != nil {
		return s.captureFail(
			e,
			fmt.Sprintf(
				"database error on %s: schemas not listed",
				instance,
			),
		)
	}

	return response.SuccessAny(v)
}
