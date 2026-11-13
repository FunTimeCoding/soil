package service

import (
	"context"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/trivago/tgo/tcontainer"
)

func (s *Service) CreateIssue(
	project string,
	issueType string,
	summary string,
	description string,
	assignee string,
	labels []string,
	fields map[string]any,
) (*issue.Issue, error) {
	raw := issue.RawStub()
	raw.Fields.Unknowns = make(tcontainer.MarshalMap)
	reporter, e := peer.Guarded(
		context.Background(),
		s.jiraGuard,
		func(context.Context) (*jira.User, error) { return s.jira.User() },
	)

	if e != nil {
		return nil, e
	}

	raw.Fields.Reporter = reporter
	raw.Fields.Project.Key = project
	raw.Fields.Type.Name = issueType
	raw.Fields.Summary = summary
	raw.Fields.Description = description
	raw.Fields.Labels = labels

	if _, f := s.applyFields(raw, fields); f != nil {
		return nil, f
	}

	created, g := peer.Guarded(
		context.Background(),
		s.jiraGuard,
		func(context.Context) (*issue.Issue, error) {
			return s.jira.CreateNative(raw)
		},
	)

	if g != nil {
		return nil, g
	}

	if assignee == "" {
		return created, nil
	}

	user, h := s.ResolveUser(assignee)

	if h != nil {
		return nil, h
	}

	if i := s.jiraGuard.Attempt(
		context.Background(),
		func(context.Context) error { return s.jira.Assign(created.Key, user) },
	); i != nil {
		return nil, i
	}

	return s.Issue(created.Key)
}
