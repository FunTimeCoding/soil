package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) SendKey(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SendKey,
) (*mcp.CallToolResult, error) {
	if len(a.Keys) == 0 {
		return response.Fail("keys is required")
	}

	interval := time.Duration(a.Interval) * time.Second

	if interval < time.Second {
		interval = time.Second
	}

	for i, key := range a.Keys {
		if i > 0 {
			time.Sleep(interval)
		}

		e := s.client.SendKey(
			a.SessionIdentifier,
			key,
		)

		if e != nil {
			return response.Fail(
				"send key %s (step %d): %v",
				key,
				i+1,
				e,
			)
		}
	}

	return response.Success(
		"sent keys: %s",
		join.Comma(a.Keys),
	)
}
