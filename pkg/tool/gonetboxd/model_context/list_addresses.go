package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listAddresses(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(generativeConstant.ParameterName)

	if f != nil {
		return response.Fail("name is required: %v", f)
	}

	result, g := s.service.DeviceAddresses(name)

	if g != nil {
		return s.captureFail(g, "addresses not retrieved")
	}

	return response.SuccessAny(convert.Addresses(result))
}
