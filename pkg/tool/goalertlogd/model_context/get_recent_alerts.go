package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) getRecentAlerts(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	now := time.Now()
	start := now.Add(-1 * time.Hour)
	end := now

	if v := r.GetString(constant.Start, ""); v != "" {
		t, e := time.Parse(time.RFC3339, v)

		if e != nil {
			return response.Fail("invalid start format: %v", e)
		}

		start = t
	}

	if v := r.GetString(constant.End, ""); v != "" {
		t, e := time.Parse(time.RFC3339, v)

		if e != nil {
			return response.Fail("invalid end format: %v", e)
		}

		end = t
	}

	records, f := s.store.ByTimeRange(start, end)

	if f != nil {
		return s.captureFail(f, "get recent alerts failed")
	}

	return response.Success(
		"Found %d alerts:\n%s",
		len(records),
		notation.MarshalIndent(records),
	)
}
