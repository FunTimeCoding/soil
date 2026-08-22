package server

import (
	"context"
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

	if e != nil {
		if isClientError(e) {
			return server.CreateIssue400JSONResponse(*clientError(e)), nil
		}

		return server.CreateIssue500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateIssue201JSONResponse(*convert.JiraIssue(created)), nil
}
