package model_context

import (
	"context"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) delete(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier := r.GetFloat(constant.Identifier, 0)

	if identifier == 0 {
		return response.Fail("id is required")
	}

	if e := s.store.Delete(uint(identifier)); e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	return response.Success("Entry %d deleted successfully", int(identifier))
}
