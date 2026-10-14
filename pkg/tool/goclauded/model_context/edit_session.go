package model_context

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
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
		return s.captureFail(e, soil.UnexpectedError)
	}

	target := q.GetString(constant.Target, "")
	sessionIdentifier, failed, f := s.targetSession(c, target)

	if failed != nil || f != nil {
		return failed, f
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
		if conflict.Is(e) {
			return response.Fail(e.Error())
		}

		return s.captureFail(e, soil.UnexpectedError)
	}

	return response.Success("Session updated")
}
