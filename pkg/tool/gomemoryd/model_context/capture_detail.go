package model_context

import (
	"database/sql"
	"errors"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
	if errors.Is(e, sql.ErrNoRows) {
		return s.captureFail(e, "not found")
	}

	return s.captureFail(e, constant.UnexpectedError)
}
