package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) tableSizes(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.TableSizes,
) (*mcp.CallToolResult, error) {
	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(e)
	}

	v, e := s.service.TableSizes(x, instance, a.Schema)

	if e != nil {
		return s.answerWith(
			e,
			fmt.Sprintf(
				"database error on %s: table sizes in %s not retrieved",
				instance,
				a.Schema,
			),
		)
	}

	return response.SuccessAny(v)
}
