package model_context

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
	if f := connection.Classify(e); f != nil {
		return s.captureFail(e, f.Error())
	}

	return s.captureFail(e, constant.UnexpectedError)
}
