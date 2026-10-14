package model_context

import (
	"context"
	"fmt"
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
		return s.captureFail(e, soil.UnexpectedError)
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

	to := q.GetString(constant.To, "")
	d := constant.Delivery(
		q.GetString(constant.DeliveryParameter, string(constant.DeliveryQueued)),
	)

	switch d {
	case constant.DeliveryQueued, constant.DeliveryImmediate, constant.DeliveryCruise:
	default:
		return response.Fail("delivery must be queued, immediate or cruise")
	}

	delivered, f := s.service.Send(c.Callsign, to, body, d)

	if f != nil {
		if not_found.Is(f) {
			return response.Fail(f.Error())
		}

		return s.captureFail(f, soil.UnexpectedError)
	}

	if to == "" {
		return response.Success("Broadcast sent")
	}

	if d == constant.DeliveryImmediate && !delivered {
		return response.Success(
			fmt.Sprintf(
				"Sent to %s, deferred to their next prompt - immediate delivery is rate limited for that session right now",
				to,
			),
		)
	}

	if d == constant.DeliveryCruise && !delivered {
		return response.Success(
			fmt.Sprintf(
				"Sent to %s, held for their cruise setting - it wakes them when the setting allows, or arrives with their next prompt",
				to,
			),
		)
	}

	return response.Success(fmt.Sprintf("Sent to %s", to))
}
