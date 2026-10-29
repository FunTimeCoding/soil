package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/types/change"
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

	c := change.New()

	if v := r.GetString(constant.Action, ""); v != "" {
		c.Action = &v
	}

	if v := r.GetString(constant.User, ""); v != "" {
		c.User = &v
	}

	if v := r.GetString(constant.System, ""); v != "" {
		c.System = &v
	}

	if v := r.GetString(constant.Service, ""); v != "" {
		c.Service = &v
	}

	if v := r.GetString(constant.Description, ""); v != "" {
		c.Description = &v
	}

	if stamp := r.GetString(constant.Timestamp, ""); stamp != "" {
		t, e := time.Parse(time.RFC3339, stamp)

		if e != nil {
			return response.Fail("invalid timestamp format: %v", e)
		}

		c.Timestamp = &t
	}

	result, f := s.service.Update(uint(identifier), c)

	if f != nil {
		return s.answer(f)
	}

	return response.Success(
		"Entry updated successfully:\n%s",
		notation.MarshalIndent(result),
	)
}
