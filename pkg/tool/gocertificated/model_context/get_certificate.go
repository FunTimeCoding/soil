package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getCertificate(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Serial,
) (*mcp.CallToolResult, error) {
	result, e := s.service.BySerial(a.Serial)

	if e != nil {
		return s.answer(e)
	}

	return response.SuccessAny(convert.Certificate(result))
}
