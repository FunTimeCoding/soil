package model_context

import (
	"context"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) add(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	action, f := r.RequireString(constant.Action)

	if f != nil {
		return response.Fail("action is required: %v", f)
	}

	user, g := r.RequireString(constant.User)

	if g != nil {
		return response.Fail("user is required: %v", g)
	}

	e := entry.New()
	e.Action = action
	e.User = user

	if y := r.GetString(constant.System, ""); y != "" {
		e.System = y
	}

	if v := r.GetString(constant.Service, ""); v != "" {
		e.Service = v
	}

	if d := r.GetString(constant.Description, ""); d != "" {
		e.Description = d
	}

	if stamp := r.GetString(constant.Timestamp, ""); stamp != "" {
		t, parseFail := time.Parse(time.RFC3339, stamp)

		if parseFail != nil {
			return response.Fail(
				"invalid timestamp format: %v",
				parseFail,
			)
		}

		e.Timestamp = t
	}

	if h := s.store.Add(e); h != nil {
		return s.captureFail(h, library.UnexpectedError)
	}

	return response.Success(
		"Entry added successfully:\n%s",
		notation.MarshalIndent(e),
	)
}
