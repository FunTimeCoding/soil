package model_context

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
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
	d := constant.Delivery(q.GetString(constant.DeliveryParameter, ""))

	switch d {
	case "", constant.DeliveryQueued, constant.DeliveryImmediate:
	default:
		return response.Fail("delivery must be immediate or queued")
	}

	held, f := s.message.Send(c.Callsign, to, body, d)

	if f != nil {
		if not_found.Is(f) {
			return response.Fail(f.Error())
		}

		return s.captureFail(x, f, soil.UnexpectedError)
	}

	return response.Success(sent(to, held, d))
}
