package model_context

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/caller"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) targetSession(
	c *caller.Caller,
	target string,
) (string, *mcp.CallToolResult, error) {
	if target == "" {
		if c.Callsign == "" || c.SessionIdentifier == "" {
			result, e := response.Fail(
				"could not resolve session - announce first or specify a target",
			)

			return "", result, e
		}

		return c.SessionIdentifier, nil, nil
	}

	r, f := s.service.ResolveSessionIdentifier(target)

	if f != nil {
		result, e := s.captureFail(f, soil.UnexpectedError)

		return "", result, e
	}

	if r.Ambiguous() {
		result, e := response.Fail(
			"ambiguous: %q matches %d sessions",
			target,
			len(r.Matches),
		)

		return "", result, e
	}

	if !r.Found() {
		result, e := response.Fail("session not found: %s", target)

		return "", result, e
	}

	return r.Identifier(), nil, nil
}
