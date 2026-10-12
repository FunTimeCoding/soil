package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createDevice(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(generativeConstant.ParameterName)

	if f != nil {
		return response.Fail("name is required: %v", f)
	}

	roleName, g := r.RequireString(constant.Role)

	if g != nil {
		return response.Fail("role is required: %v", g)
	}

	typeName, h := r.RequireString(constant.Type)

	if h != nil {
		return response.Fail("type is required: %v", h)
	}

	siteName, i := r.RequireString(constant.Site)

	if i != nil {
		return response.Fail("site is required: %v", i)
	}

	result, e := s.service.CreateDevice(
		name,
		roleName,
		typeName,
		siteName,
		r.GetString(constant.Tenant, ""),
		r.GetStringSlice(constant.Tags, nil),
	)

	if e != nil {
		return s.serviceFail(e, "device not created")
	}

	return response.SuccessAny(convert.Device(result))
}
