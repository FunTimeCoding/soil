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
	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	schema := a.Schema

	if schema == "" {
		schema = "public"
	}

	v, e := s.service.TableSizes(x, instance, schema)

	if e != nil {
		return s.captureFail(
			e,
			fmt.Sprintf(
				"database error on %s: table sizes in %s not retrieved",
				instance,
				schema,
			),
		)
	}

	return response.SuccessAny(v)
}
