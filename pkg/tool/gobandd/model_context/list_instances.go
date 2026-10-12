package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/model_context/instance_response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listInstances(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListInstances,
) (*mcp.CallToolResult, error) {
	active := s.activeInstanceName(x)
	var result []*instance_response.Response

	for _, i := range s.service.Instances() {
		result = append(
			result,
			instance_response.New(i.Name, i.Host, i.Name == active),
		)
	}

	return response.SuccessAny(result)
}
