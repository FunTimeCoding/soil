package model_context

import (
	"context"
	"errors"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/argument/edit_session"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) editSession(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c, e := s.resolveCaller(x, constant.EditSession)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	target := q.GetString(constant.Target, "")
	var sessionIdentifier string

	if target != "" {
		r, f := s.service.ResolveSessionIdentifier(target)

		if f != nil {
			return s.captureFail(f, library.UnexpectedError)
		}

		if r.Ambiguous() {
			return response.Fail(
				"ambiguous: %q matches %d sessions",
				target,
				len(r.Matches),
			)
		}

		if !r.Found() {
			return response.Fail("session not found: %s", target)
		}

		sessionIdentifier = r.Identifier()
	} else if c.Callsign != "" {
		sessionIdentifier = c.SessionIdentifier
	}

	if sessionIdentifier == "" {
		return response.Fail(
			"could not resolve session - announce first or specify a target",
		)
	}

	a := edit_session.New()
	hasChange := false

	if v := q.GetString(constant.Alias, ""); v != "" {
		a.Alias = &v
		hasChange = true
	}

	if v := q.GetString(constant.Description, ""); v != "" {
		a.Description = &v
		hasChange = true
	}

	if v := q.GetString(constant.Slug, ""); v != "" {
		a.Slug = &v
		hasChange = true
	}

	if v := q.GetString(constant.Topic, ""); v != "" {
		a.Topic = &v
		hasChange = true
	}

	if v := q.GetString(constant.Files, ""); v != "" {
		a.Files = &v
		hasChange = true
	}

	if !hasChange {
		return response.Fail("at least one field required")
	}

	if e := s.service.EditSession(sessionIdentifier, a); e != nil {
		if errors.Is(e, constant.ErrorAliasCollision) {
			return response.Fail(e.Error())
		}

		return s.captureFail(e, library.UnexpectedError)
	}

	return response.Success("Session updated")
}
