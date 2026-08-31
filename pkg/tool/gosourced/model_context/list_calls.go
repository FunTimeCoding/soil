package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listCalls(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListCalls,
) (*mcp.CallToolResult, error) {
	if a.Region == "" {
		return response.Fail("region is required")
	}

	if a.Limit < 0 {
		return response.Fail("limit must not be negative")
	}

	directory, e := s.resolveDirectory(x)

	if e != nil {
		return response.Fail("%s", e)
	}

	limit := int(a.Limit)

	if limit == 0 {
		limit = 100
	}

	r, inventory, f := s.service.ListCalls(directory, a.Region, limit)

	if f != nil {
		return s.captureFail(f, constant.UnexpectedError)
	}

	if inventory == nil {
		return response.Fail("%s", formatConcerns(r.Entries))
	}

	return response.SuccessAny(inventory)
}
