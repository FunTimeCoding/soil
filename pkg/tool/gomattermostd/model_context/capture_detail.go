package model_context

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mattermost/mattermost/server/public/model"
)

func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
	if f, okay := errors.AsType[*model.AppError](e); okay {
		if f.Message != "" {
			return s.captureFail(e, f.Message)
		}
	}

	return s.captureFail(e, constant.UnexpectedError)
}
