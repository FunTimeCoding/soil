package model_context

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/refusal"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) send(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c, e := s.resolveCaller(x, constant.Send)

	if e != nil {
		return s.captureFail(x, e, soil.UnexpectedError)
	}

	if c.Callsign == "" {
		return response.Fail(
			"unknown session - announce first to bind your identity",
		)
	}

	body, e := q.RequireString(constant.Body)

	if e != nil {
		return response.Fail("body is required: %v", e)
	}

	to := q.GetStringSlice(constant.To, nil)
	result, f := s.message.Send(
		c.Callsign,
		to,
		body,
		constant.Delivery(q.GetString(constant.DeliveryParameter, "")),
	)

	if f != nil {
		if not_found.Is(f) || refusal.Is(f) {
			return response.Fail(f.Error())
		}

		return s.captureFail(x, f, soil.UnexpectedError)
	}

	return response.Success(sent(to, result))
}
