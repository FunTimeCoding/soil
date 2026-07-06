package model_context

import (
	"context"
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) update(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c, e := s.resolveCaller(x, constant.Update)

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

	topic, e := q.RequireString(constant.Topic)

	if e != nil {
		return response.Fail("topic is required: %v", e)
	}

	var files []string

	if raw, okay := q.GetArguments()[constant.Files]; okay {
		if list, okay1 := raw.([]any); okay1 {
			for _, item := range list {
				if str, okay2 := item.(string); okay2 {
					files = append(files, str)
				}
			}
		}
	}

	if f := s.service.Update(
		c.SessionIdentifier,
		c.Callsign,
		topic,
		message,
		join.NewLine(files),
	); f != nil {
		return s.captureFail(f, library.UnexpectedError)
	}

	return response.Success(
		fmt.Sprintf("Updated %s: %s", topic, message),
	)
}
