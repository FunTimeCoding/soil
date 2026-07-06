package model_context

import (
	"context"
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) historyCount(
	x context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	_, e := s.resolveCaller(x, constant.HistoryCount)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	count, e := s.service.CountEvents()

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	return response.Success(fmt.Sprintf("%d events", count))
}
