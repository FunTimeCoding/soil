package model_context

import (
	"context"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) moment(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c, e := s.resolveCaller(x, constant.Moment)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	if c.Callsign == "" {
		return response.Fail("unknown session - announce first to bind your identity")
	}

	line, e := q.RequireString(constant.Line)

	if e != nil {
		return response.Fail("line is required: %v", e)
	}

	if f := s.service.Moment(
		c.SessionIdentifier,
		c.Callsign,
		line,
	); f != nil {
		return s.captureFail(f, library.UnexpectedError)
	}

	return response.Success("moment captured")
}
