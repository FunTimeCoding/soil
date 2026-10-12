package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createCluster(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(generativeConstant.ParameterName)

	if f != nil {
		return response.Fail("name is required: %v", f)
	}

	typeName, g := r.RequireString(constant.Type)

	if g != nil {
		return response.Fail("type is required: %v", g)
	}

	siteName, h := r.RequireString(constant.Site)

	if h != nil {
		return response.Fail("site is required: %v", h)
	}

	result, e := s.service.CreateCluster(name, typeName, siteName)

	if e != nil {
		return s.serviceFail(e, "cluster not created")
	}

	return response.SuccessAny(convert.Cluster(result))
}
