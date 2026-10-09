package model_context

import (
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	reacherConstant "github.com/funtimecoding/soil/pkg/reacher/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
	detail := classify.Message(e, constant.RequestFailed)

	if edge := s.reacher.Observe(s.host, e); edge != nil {
		s.reporter.CaptureWithContext(
			e,
			reacherConstant.ContextKey,
			edge.Context(),
		)

		return response.Fail(detail)
	}

	if down, h := s.reacher.Down(s.host); down {
		return response.Fail(
			"%s unreachable since %s: %s",
			h.Name,
			h.Since.Format(constant.SinceFormat),
			h.Reason,
		)
	}

	if !classify.Reportable(e) {
		return response.Fail(detail)
	}

	return s.captureFail(e, detail)
}
