package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/trivago/tgo/tcontainer"
)

func (s *Service) UpdateIssue(
	key string,
	summary string,
	description string,
	assignee string,
	reporter string,
	labels []string,
	fields map[string]any,
) (*IssueUpdate, error) {
	if summary == "" &&
		description == "" &&
		assignee == "" &&
		reporter == "" &&
		labels == nil &&
		len(fields) == 0 {
		return nil, not_selected.Format("no fields to update")
	}

	before, e := s.jira.Issue(key)

	if e != nil {
		return nil, fmt.Errorf("issue not found: %w", e)
	}

	raw := issue.Raw(key)
	raw.Fields.Unknowns = make(tcontainer.MarshalMap)

	if summary != "" {
		raw.Fields.Summary = summary
	}

	if description != "" {
		raw.Fields.Description = description
	}

	if reporter != "" {
		user, f := s.ResolveUser(reporter)

		if f != nil {
			return nil, f
		}

		raw.Fields.Reporter = user
	}

	if labels != nil {
		raw.Fields.Labels = labels
	}

	customFieldNames, g := s.applyFields(raw, fields)

	if g != nil {
		return nil, g
	}

	hasFieldChanges := summary != "" ||
		description != "" ||
		reporter != "" ||
		labels != nil ||
		len(fields) > 0

	if hasFieldChanges {
		if h := s.jira.UpdateNative(raw); h != nil {
			return nil, h
		}
	}

	if assignee != "" {
		user, i := s.ResolveUser(assignee)

		if i != nil {
			return nil, i
		}

		if j := s.jira.Assign(key, user); j != nil {
			return nil, j
		}
	}

	after, k := s.jira.Issue(key)

	if k != nil {
		return nil, fmt.Errorf("issue updated but retrieval failed: %w", k)
	}

	return &IssueUpdate{
		Before:           before,
		After:            after,
		CustomFieldNames: customFieldNames,
	}, nil
}
