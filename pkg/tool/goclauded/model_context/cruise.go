package model_context

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/refusal"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) cruise(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c, e := s.resolveCaller(x, constant.CruiseTool)

	if e != nil {
		return s.captureFail(x, e, soil.UnexpectedError)
	}

	sessionIdentifier, failed, f := s.targetSession(
		x,
		c,
		q.GetString(constant.Target, ""),
	)

	if failed != nil || f != nil {
		return failed, f
	}

	if mode := q.GetString(constant.Mode, ""); mode != "" {
		if g := s.message.SetCruise(
			sessionIdentifier,
			constant.Cruise(mode),
			time.Duration(q.GetFloat(constant.Pace, 0)*float64(time.Minute)),
		); g != nil {
			if refusal.Is(g) {
				return response.Fail(g.Error())
			}

			return s.captureFail(x, g, soil.UnexpectedError)
		}
	}

	setting, h := s.message.CruiseSetting(sessionIdentifier)

	if h != nil {
		return s.captureFail(x, h, soil.UnexpectedError)
	}

	return response.Success(cruiseLine(setting))
}
