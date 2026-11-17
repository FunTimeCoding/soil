package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateVirtualMachine(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateVirtualMachine,
) (*mcp.CallToolResult, error) {
	if a.Name == "" || a.Cluster == "" {
		return response.Fail("name and cluster are required")
	}

	result, e := s.service.CreateVirtualMachine(a.Name, a.Cluster)

	if e != nil {
		return s.answerMessage(x, e, "virtual machine not created")
	}

	return response.SuccessAny(convert.VirtualMachine(result))
}
