package model_context

import (
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	peerConstant "github.com/funtimecoding/soil/pkg/peer/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) answer(e error) (*mcp.CallToolResult, error) {
	if edge := s.reacher.Observe(s.host, e); edge != nil {
		s.reporter.CaptureWithContext(
			e,
			peerConstant.ContextKey,
			edge.Context(),
		)

		return response.Fail("%s", e)
	}

	if down, h := s.reacher.Down(s.host); down {
		return response.Fail(
			"%s unreachable since %s: %s",
			h.Name,
			h.Since.Format(constant.SinceFormat),
			h.Reason,
		)
	}

	return s.answerWith(e, "")
}
