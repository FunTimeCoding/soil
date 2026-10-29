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

func (s *Server) UpdateIssue(
	x context.Context,
	r server.UpdateIssueRequestObject,
) (server.UpdateIssueResponseObject, error) {
	summary := ""
	description := ""
	assignee := ""
	reporter := ""
	noDiff := false
	var labels []string
	var fields map[string]any

	if r.Body.Summary != nil {
		summary = *r.Body.Summary
	}

	if r.Body.Description != nil {
		description = *r.Body.Description
	}

	if r.Body.Assignee != nil {
		assignee = *r.Body.Assignee
	}

	if r.Body.Reporter != nil {
		reporter = *r.Body.Reporter
	}

	if r.Body.NoDiff != nil {
		noDiff = *r.Body.NoDiff
	}

	if r.Body.Labels != nil {
		labels = *r.Body.Labels
	}

	if r.Body.AdditionalFields != nil {
		fields = *r.Body.AdditionalFields
	}

	result, e := s.service.UpdateIssue(
		r.Key,
		summary,
		description,
		assignee,
		reporter,
		labels,
		fields,
	)

	if validation.Is(e) || ambiguous.Is(e) || not_selected.Is(e) {
		return server.UpdateIssue400JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.UpdateIssue403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.UpdateIssue404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.UpdateIssue503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.UpdateIssue500JSONResponse(*s.answer(e)), nil
	}

	diff := convert.JiraIssueDiff(
		result.Before,
		result.After,
		noDiff,
		result.CustomFieldNames,
	)
	response := server.IssueUpdateResult{Issue: *diff.Issue}

	if len(diff.Changes) > 0 {
		changes := make([]server.FieldChange, 0, len(diff.Changes))

		for _, c := range diff.Changes {
			changes = append(
				changes,
				server.FieldChange{
					Field:  c.Field,
					Before: c.Before,
					After:  c.After,
				},
			)
		}

		response.Changes = &changes
	}

	return server.UpdateIssue200JSONResponse(response), nil
}
