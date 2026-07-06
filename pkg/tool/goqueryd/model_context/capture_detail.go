package model_context

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/mark3labs/mcp-go/mcp"
	"gorm.io/gorm"
)

func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
	if errors.Is(e, gorm.ErrRecordNotFound) {
		return s.captureFail(e, e.Error())
	}

	return s.captureFail(e, constant.UnexpectedError)
}
