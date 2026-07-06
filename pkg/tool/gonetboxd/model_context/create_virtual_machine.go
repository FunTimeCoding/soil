package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createVirtualMachine(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(parameter.Name)

	if f != nil {
		return response.Fail("name is required: %v", f)
	}

	clusterName, g := r.RequireString(constant.Cluster)

	if g != nil {
		return response.Fail("cluster is required: %v", g)
	}

	cl, h := s.client.ClusterByName(clusterName)

	if h != nil {
		return s.captureFail(h, "cluster not found")
	}

	result, i := s.client.CreateVirtualMachine(name, cl)

	if i != nil {
		return s.captureFail(i, "virtual machine not created")
	}

	return response.SuccessAny(convert.VirtualMachine(result))
}
