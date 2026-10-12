package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createTunnel(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, e := r.RequireString(generativeConstant.ParameterName)

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

	t, h := s.service.CreateTunnel(name, encapsulation, groupName)

	if h != nil {
		return s.serviceFail(h, "tunnel not created")
	}

	return response.SuccessAny(convert.Tunnel(t))
}
