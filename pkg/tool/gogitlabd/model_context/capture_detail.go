package model_context

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
	if f, okay := errors.AsType[*gitlab.ErrorResponse](e); okay {
		if f.Message != "" {
			return s.captureFail(e, f.Message)
		}
	}

	return s.captureFail(e, constant.UnexpectedError)
}
