package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateCluster(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateCluster,
) (*mcp.CallToolResult, error) {
	if a.Name == "" || a.Type == "" || a.Site == "" {
		return response.Fail("name, type and site are required")
	}

	result, e := s.service.CreateCluster(a.Name, a.Type, a.Site)

	if e != nil {
		return s.answerMessage(x, e, "cluster not created")
	}

	return response.SuccessAny(convert.Cluster(result))
}
