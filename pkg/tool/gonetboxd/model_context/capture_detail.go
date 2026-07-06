package model_context

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/common"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
	if m, okay := common.ExtractMessage(e); okay {
		return s.captureFail(e, m)
	}

	return s.captureFail(e, constant.UnexpectedError)
}
