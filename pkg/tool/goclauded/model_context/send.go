package model_context

import (
	"context"
	"errors"
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
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
		return s.captureFail(e, library.UnexpectedError)
	}

	if c.Callsign == "" {
		return response.Fail("unknown session - announce first to bind your identity")
	}

	body, e := q.RequireString(constant.Body)

	if e != nil {
		return response.Fail("body is required: %v", e)
	}

	to := q.GetString(constant.To, "")

	if f := s.service.Send(c.Callsign, to, body); f != nil {
		if errors.Is(f, constant.ErrorCallsignNotFound) {
			return response.Fail(f.Error())
		}

		return s.captureFail(f, library.UnexpectedError)
	}

	if to != "" {
		return response.Success(fmt.Sprintf("Sent to %s", to))
	}

	return response.Success("Broadcast sent")
}
