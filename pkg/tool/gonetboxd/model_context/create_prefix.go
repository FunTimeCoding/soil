package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createPrefix(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	cidr, e := r.RequireString(constant.Prefix)

	if e != nil {
		return response.Fail("prefix is required: %v", e)
	}

	p, f := s.service.CreatePrefix(
		cidr,
		r.GetString(constant.Site, ""),
		r.GetString(constant.Description, ""),
	)

	if f != nil {
		return s.serviceFail(f, "prefix not created")
	}

	return response.SuccessAny(convert.Prefix(p))
}
