package model_context

import (
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	peerConstant "github.com/funtimecoding/soil/pkg/peer/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) succeed(v any) (*mcp.CallToolResult, error) {
	if edge := s.reacher.Observe(s.host, nil); edge != nil {
		s.reporter.CaptureWithContext(
			recovered(edge),
			peerConstant.ContextKey,
			edge.Context(),
		)
	}

	return response.SuccessAny(v)
}
