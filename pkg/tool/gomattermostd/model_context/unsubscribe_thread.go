package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) UnsubscribeThread(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.UnsubscribeThread,
) (*mcp.CallToolResult, error) {
	if a.Root == "" {
		return response.Fail("root is required")
	}

	if a.Callsign == "" {
		return response.Fail("callsign is required")
	}

	count, e := s.store.Delete(a.Callsign, a.Root)

	if e != nil {
		return s.captureDetail(e)
	}

	if count == 0 {
		return response.Success("no such subscription")
	}

	remaining, f := s.store.ByRoot(a.Root)

	if f != nil {
		return s.captureDetail(f)
	}

	if len(remaining) == 0 {
		s.indexer.Forget(a.Root)
	}

	return response.Success("unsubscribed")
}
