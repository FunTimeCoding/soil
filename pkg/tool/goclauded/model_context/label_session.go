package model_context

import (
	"context"
	"fmt"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
	"slices"
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

	_, set := q.GetArguments()[constant.Value]
	add := q.GetString(constant.LabelAdd, "")
	remove := q.GetString(constant.LabelRemove, "")

	if !set && add == "" && remove == "" {
		value, g := s.label.GetLabel(sessionIdentifier, key)

		if g != nil {
			return s.captureFail(x, g, soil.UnexpectedError)
		}

		if value == "" {
			return response.Success("no label %s", key)
		}

		return response.Success("%s=%s", key, value)
	}

	writes := add != "" || q.GetString(constant.Value, "") != ""

	if writes && slices.Contains(constant.ReservedLabelKeys, key) {
		return response.Fail(
			"reserved key %q - session metadata belongs in edit_session",
			key,
		)
	}

	change, h := s.changeLabel(sessionIdentifier, c.Callsign, target, key, q)

	if h != nil {
		return s.captureFail(x, h, soil.UnexpectedError)
	}

	return response.Success(
		fmt.Sprintf("%s on %s", change, displayName(c, target)),
	)
}
