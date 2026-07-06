package model_context

import (
	"context"
	"errors"
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) announce(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, e := q.RequireString(constant.SessionName)

	if e != nil {
		return response.Fail("name is required: %v", e)
	}

	topic, e := q.RequireString(constant.Topic)

	if e != nil {
		return response.Fail("topic is required: %v", e)
	}

	var files []string

	if raw, okay := q.GetArguments()[constant.Files]; okay {
		if list, okay := raw.([]any); okay {
			for _, item := range list {
				if str, okay := item.(string); okay {
					files = append(files, str)
				}
			}
		}
	}

	if cs := server.ClientSessionFromContext(x); cs != nil {
		if f := s.service.BindModelContextSession(
			name,
			cs.SessionID(),
		); f != nil {
			if errors.Is(f, constant.ErrorCallsignNotFound) {
				return response.Fail(f.Error())
			}

			return s.captureFail(f, library.UnexpectedError)
		}
	}

	c, g := s.resolveCaller(x, constant.Announce)

	if g != nil {
		return s.captureFail(g, library.UnexpectedError)
	}

	if f := s.service.Announce(
		c.SessionIdentifier,
		name,
		topic,
		join.NewLine(files),
	); f != nil {
		if errors.Is(f, constant.ErrorCallsignNotFound) {
			return response.Fail(f.Error())
		}

		return s.captureFail(f, library.UnexpectedError)
	}

	return response.Success(
		fmt.Sprintf(
			"Announced as %s: %s",
			name,
			topic,
		),
	)
}
