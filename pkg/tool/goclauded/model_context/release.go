package model_context

import (
	"context"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) release(
	x context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c, e := s.resolveCaller(x, constant.Release)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	if c.Callsign == "" {
		return response.Fail("unknown session - announce first to bind your identity")
	}

	if e := s.service.Release(c.SessionIdentifier, c.Callsign); e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	return response.Success("Released")
}
