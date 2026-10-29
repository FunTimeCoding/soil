package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) CreateIssue(
	x context.Context,
	r server.CreateIssueRequestObject,
) (server.CreateIssueResponseObject, error) {
	description := ""
	assignee := ""
	var labels []string
	var fields map[string]any

	if r.Body.Description != nil {
		description = *r.Body.Description
	}

	if r.Body.Assignee != nil {
		assignee = *r.Body.Assignee
	}

	if r.Body.Labels != nil {
		labels = *r.Body.Labels
	}

	if r.Body.AdditionalFields != nil {
		fields = *r.Body.AdditionalFields
	}

	created, e := s.service.CreateIssue(
		r.Body.Project,
		r.Body.IssueType,
		r.Body.Summary,
		description,
		assignee,
		labels,
		fields,
	)

	if validation.Is(e) || ambiguous.Is(e) || not_selected.Is(e) {
		return server.CreateIssue400JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.CreateIssue403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.CreateIssue404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.CreateIssue503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreateIssue500JSONResponse(*s.answer(e)), nil
	}

	return server.CreateIssue201JSONResponse(*convert.JiraIssue(created)), nil
}
