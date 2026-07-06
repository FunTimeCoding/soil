package model_context

import (
	"context"
	"errors"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) update(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier := r.GetFloat(constant.Identifier, 0)

	if identifier == 0 {
		return response.Fail("id is required")
	}

	e, f := s.store.Get(uint(identifier))

	if f != nil {
		if errors.Is(f, not_found.Sentinel) {
			return response.Fail("%s", f)
		}

		return s.captureFail(f, library.UnexpectedError)
	}

	if v := r.GetString(constant.Action, ""); v != "" {
		e.Action = v
	}

	if v := r.GetString(constant.User, ""); v != "" {
		e.User = v
	}

	if v := r.GetString(constant.System, ""); v != "" {
		e.System = v
	}

	if v := r.GetString(constant.Service, ""); v != "" {
		e.Service = v
	}

	if v := r.GetString(constant.Description, ""); v != "" {
		e.Description = v
	}

	if stamp := r.GetString(constant.Timestamp, ""); stamp != "" {
		t, g := time.Parse(time.RFC3339, stamp)

		if g != nil {
			return response.Fail(
				"invalid timestamp format: %v",
				g,
			)
		}

		e.Timestamp = t
	}

	if g := s.store.Update(e); g != nil {
		return s.captureFail(g, library.UnexpectedError)
	}

	return response.Success(
		"Entry updated successfully:\n%s",
		notation.MarshalIndent(e),
	)
}
