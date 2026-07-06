package model_context

import (
	"context"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) list(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	filter := store.NewFilter()

	if y := r.GetString(constant.System, ""); y != "" {
		filter.System = y
	}

	if v := r.GetString(constant.Service, ""); v != "" {
		filter.Service = v
	}

	if u := r.GetString(constant.User, ""); u != "" {
		filter.User = u
	}

	if since := r.GetString(constant.Since, ""); since != "" {
		i, parseFail := time.Parse(time.RFC3339, since)

		if parseFail != nil {
			return response.Fail(
				"invalid since format: %v",
				parseFail,
			)
		}

		filter.Since = i
	}

	if until := r.GetString(constant.Until, ""); until != "" {
		u, parseFail := time.Parse(time.RFC3339, until)

		if parseFail != nil {
			return response.Fail(
				"invalid until format: %v",
				parseFail,
			)
		}

		filter.Until = u
	}

	if l := r.GetFloat(parameter.Limit, 0); l > 0 {
		filter.Limit = int(l)
	}

	entries, e := s.store.List(filter)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	return response.Success(
		"Found %d entries:\n%s",
		len(entries),
		notation.MarshalIndent(entries),
	)
}
