package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListRemediation(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListRemediation,
) (*mcp.CallToolResult, error) {
	m := constant.CriticalScore

	if a.Score != nil {
		m = *a.Score
	}

	result, e := s.remediation.List(m)

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(convert.RemediationList(result))
}
