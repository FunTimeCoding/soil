package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) pendingPublication(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ struct{},
) (*mcp.CallToolResult, error) {
	result, e := s.service.Pending()

	if e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	return response.SuccessAny(convert.Changes(result))
}
