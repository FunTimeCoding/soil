package model_context

import (
	"context"
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) label(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c, e := s.resolveCaller(x, constant.Label)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	key, e := q.RequireString(constant.Key)

	if e != nil {
		return response.Fail("key is required: %v", e)
	}

	target := q.GetString(constant.Target, "")
	var sessionIdentifier string
	var targetName string

	if target != "" {
		r, f := s.service.ResolveSessionIdentifier(target)

		if f != nil {
			return s.captureFail(f, library.UnexpectedError)
		}

		if r.Ambiguous() {
			return response.Fail(
				"ambiguous: %q matches %d sessions",
				target,
				len(r.Matches),
			)
		}

		if !r.Found() {
			return response.Fail("session not found: %s", target)
		}

		sessionIdentifier = r.Identifier()
		targetName = target
	} else if c.Callsign != "" {
		sessionIdentifier = c.SessionIdentifier
	}

	if sessionIdentifier == "" {
		return response.Fail(
			"could not resolve session - announce first or specify a target",
		)
	}

	arguments := q.GetArguments()
	raw, hasValue := arguments[constant.Value]

	if !hasValue {
		value, f := s.service.GetLabel(sessionIdentifier, key)

		if f != nil {
			return s.captureFail(f, library.UnexpectedError)
		}

		if value == "" {
			return response.Success("no label %s", key)
		}

		return response.Success("%s=%s", key, value)
	}

	value, _ := raw.(string)

	if value == "" {
		change, g := s.service.DeleteLabel(
			sessionIdentifier,
			c.Callsign,
			targetName,
			key,
		)

		if g != nil {
			return s.captureFail(g, library.UnexpectedError)
		}

		return response.Success(change)
	}

	for _, reserved := range constant.ReservedLabelKeys {
		if key == reserved {
			return response.Fail(
				"reserved key %q - session metadata belongs in edit_session",
				key,
			)
		}
	}

	change, h := s.service.SetLabel(
		sessionIdentifier,
		c.Callsign,
		targetName,
		key,
		value,
	)

	if h != nil {
		return s.captureFail(h, library.UnexpectedError)
	}

	return response.Success(
		fmt.Sprintf("%s on %s", change, displayName(c, target)),
	)
}
