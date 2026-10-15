package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deleteCustomField(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, e := r.RequireString(generativeConstant.ParameterName)

	if e != nil {
		return response.Fail("name is required: %v", e)
	}

	f := s.client.DeleteCustomField(name)

	if not_found.Is(f) {
		return response.Fail("%v", f)
	}

	if f != nil {
		return s.captureFail(f, "custom field not deleted")
	}

	return response.Success("custom field deleted")
}
