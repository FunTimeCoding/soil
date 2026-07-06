package model_context

import (
	"context"
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) complete(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c, e := s.resolveCaller(x, constant.Complete)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	if c.Callsign == "" {
		return response.Fail("unknown session - announce first to bind your identity")
	}

	message, e := q.RequireString(constant.Message)

	if e != nil {
		return response.Fail("message is required: %v", e)
	}

	topic, f := s.service.CompleteTask(c.Callsign)

	if f != nil {
		return s.captureFail(f, library.UnexpectedError)
	}

	if topic == "" {
		topic = q.GetString(constant.Topic, "")
	}

	if topic == "" {
		return response.Fail("no active topic - provide a topic parameter or announce first")
	}

	if f := s.service.Complete(
		c.SessionIdentifier,
		c.Callsign,
		topic,
		message,
	); f != nil {
		return s.captureFail(f, library.UnexpectedError)
	}

	return response.Success(
		fmt.Sprintf("Completed %s: %s", topic, message),
	)
}
