package model_context

import (
	"context"
	"fmt"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) labelSession(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c, e := s.resolveCaller(x, constant.Label)

	if e != nil {
		return s.captureFail(x, e, soil.UnexpectedError)
	}

	key, e := q.RequireString(constant.Key)

	if e != nil {
		return response.Fail("key is required: %v", e)
	}

	target := q.GetString(constant.Target, "")
	sessionIdentifier, failed, f := s.targetSession(x, c, target)

	if failed != nil || f != nil {
		return failed, f
	}

	targetName := target
	arguments := q.GetArguments()
	raw, hasValue := arguments[constant.Value]

	if !hasValue {
		value, f := s.label.GetLabel(sessionIdentifier, key)

		if f != nil {
			return s.captureFail(x, f, soil.UnexpectedError)
		}

		if value == "" {
			return response.Success("no label %s", key)
		}

		return response.Success("%s=%s", key, value)
	}

	value, _ := raw.(string)

	if value == "" {
		change, g := s.label.DeleteLabel(
			sessionIdentifier,
			c.Callsign,
			targetName,
			key,
		)

		if g != nil {
			return s.captureFail(x, g, soil.UnexpectedError)
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

	change, h := s.label.SetLabel(
		sessionIdentifier,
		c.Callsign,
		targetName,
		key,
		value,
	)

	if h != nil {
		return s.captureFail(x, h, soil.UnexpectedError)
	}

	return response.Success(
		fmt.Sprintf("%s on %s", change, displayName(c, target)),
	)
}
