package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/time/window"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) getRecentAlerts(
	x context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	w := window.Whole()

	if v := r.GetString(constant.Start, ""); v != "" {
		t, e := time.Parse(time.RFC3339, v)

		if e != nil {
			return response.Fail("invalid start format: %v", e)
		}

		w.WithStart(t)
	}

	if v := r.GetString(constant.End, ""); v != "" {
		t, e := time.Parse(time.RFC3339, v)

		if e != nil {
			return response.Fail("invalid end format: %v", e)
		}

		w.WithEnd(t)
	}

	records, f := s.service.RecentAlerts(w)

	if f != nil {
		return s.answer(x, f)
	}

	return response.Success(
		"Found %d alerts:\n%s",
		len(records),
		notation.MarshalIndent(records),
	)
}
