package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateClusterType(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Name,
) (*mcp.CallToolResult, error) {
	if a.Name == "" {
		return response.Fail("name is required")
	}

	result, e := s.service.CreateClusterType(a.Name)

	if e != nil {
		return s.answerMessage(x, e, "cluster type not created")
	}

	return response.SuccessAny(convert.ClusterType(result))
}
