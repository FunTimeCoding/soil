package model_context

import (
	"context"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) linkIssues(
	c context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(generative.ParameterKey)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	target, g := r.RequireString(constant.TargetKey)

	if g != nil {
		return response.Fail("target_key is required: %v", g)
	}

	linkType := r.GetString(constant.LinkType, "Relates")

	if h := s.service.LinkIssues(key, target, linkType); h != nil {
		return response.Fail("%s", h)
	}

	return response.Success("%s %s %s", key, linkType, target)
}
