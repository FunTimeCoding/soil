package model_context

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
	if d, okay := errors.AsType[*detail_error.Detail](e); okay {
		return s.captureFail(e, d.Detail)
	}

	detail := classify.Message(e, constant.RequestFailed)

	if !classify.Reportable(e) {
		return response.Fail(detail)
	}

	return s.captureFail(e, detail)
}
