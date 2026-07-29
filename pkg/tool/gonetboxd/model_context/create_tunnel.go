package model_context

import (
	"context"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createTunnel(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, e := r.RequireString(generative.ParameterName)

	if e != nil {
		return response.Fail("name is required: %v", e)
	}

	encapsulation, f := r.RequireString(constant.Encapsulation)

	if f != nil {
		return response.Fail("encapsulation is required: %v", f)
	}

	groupName, g := r.RequireString(constant.Group)

	if g != nil {
		return response.Fail("group is required: %v", g)
	}

	group, h := s.client.TunnelGroupByName(groupName)

	if h != nil {
		return s.captureDetail(h)
	}

	t, i := s.client.CreateTunnel(name, encapsulation, group)

	if i != nil {
		return s.captureDetail(i)
	}

	return response.SuccessAny(convert.Tunnel(t))
}
