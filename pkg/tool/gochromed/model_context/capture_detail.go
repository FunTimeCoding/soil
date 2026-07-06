package model_context

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
	if errors.Is(e, context.DeadlineExceeded) {
		return s.captureFail(
			e,
			"tab timed out - it may be sleeping; activate it in the browser first",
		)
	}

	return s.captureFail(e, constant.UnexpectedError)
}
