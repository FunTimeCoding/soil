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
	if e := s.service.SendKeys(
		a.SessionIdentifier,
		a.Keys,
		time.Duration(a.Interval)*time.Second,
	); e != nil {
		return s.answer(e)
	}

	return response.Success("sent keys: %s", join.Comma(a.Keys))
}
