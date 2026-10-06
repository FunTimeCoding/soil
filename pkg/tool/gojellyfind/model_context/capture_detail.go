package model_context

import (
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
	detail := classify.Message(e, constant.RequestFailed)

	if !classify.Reportable(e) {
		return response.Fail(detail)
	}

	return s.captureFail(e, detail)
}
