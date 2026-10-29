package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deleteConfigurationContext(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, e := r.RequireString(generativeConstant.ParameterName)

	if e != nil {
		return response.Fail("name is required: %v", e)
	}

	f := s.service.DeleteConfigurationContext(name)

	if not_found.Is(f) {
		return response.Fail("%v", f)
	}

	if f != nil {
		return s.captureFail(f, "configuration context not deleted")
	}

	return response.Success("configuration context deleted")
}
