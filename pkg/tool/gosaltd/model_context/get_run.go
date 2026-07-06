package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getRun(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier := uint(r.GetInt(constant.Identifier, 0))

	if identifier == 0 {
		return response.Fail("id is required")
	}

	result, e := s.store.Find(identifier)

	if e != nil {
		return s.captureFail(e, constant.RunLookupFailed)
	}

	return response.SuccessAny(result)
}
