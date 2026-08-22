package server

import (
	"context"
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

	if e != nil {
		if isClientError(e) {
			return server.UpdateIssue400JSONResponse(*clientError(e)), nil
		}

		return server.UpdateIssue500JSONResponse(*s.captureDetail(e)), nil
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
