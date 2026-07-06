package model_context

import (
	"context"
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) status(
	x context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c, e := s.resolveCaller(x, constant.Status)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	if c.Callsign == "" {
		return response.Fail(
			"unknown session - announce first to bind your identity",
		)
	}

	d, f := s.service.SessionByCallsign(c.Callsign)

	if f != nil {
		return s.captureFail(f, library.UnexpectedError)
	}

	if d == nil {
		return response.Fail("session not found for %s", c.Callsign)
	}

	lines := []string{fmt.Sprintf("Name: %s", d.Name)}

	if d.Topic != "" {
		lines = append(lines, fmt.Sprintf("Topic: %s", d.Topic))
	} else {
		lines = append(lines, "Topic: (none)")
	}

	if d.Files != "" {
		lines = append(lines, fmt.Sprintf("Files: %s", d.Files))
	}

	return response.Success(join.NewLine(lines))
}
