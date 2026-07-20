package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getStatus(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	totalRecords, e := s.store.Count()

	if e != nil {
		return s.captureFail(e, "get status failed")
	}

	lastPoll := s.worker.LastPoll()

	if lastPoll.IsZero() {
		return response.Success(
			"Total records: %d\nLast poll: never",
			totalRecords,
		)
	}

	return response.Success(
		"Total records: %d\nLast poll: %s",
		totalRecords,
		lastPoll.Format(constant.DateFormat),
	)
}
