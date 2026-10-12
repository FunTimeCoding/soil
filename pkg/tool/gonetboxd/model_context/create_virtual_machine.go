package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createVirtualMachine(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(generativeConstant.ParameterName)

	if f != nil {
		return response.Fail("name is required: %v", f)
	}

	clusterName, g := r.RequireString(constant.Cluster)

	if g != nil {
		return response.Fail("cluster is required: %v", g)
	}

	result, e := s.service.CreateVirtualMachine(name, clusterName)

	if e != nil {
		return s.serviceFail(e, "virtual machine not created")
	}

	return response.SuccessAny(convert.VirtualMachine(result))
}
