package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

type IssueClient interface {
	SearchIssues(
		x context.Context,
		query string,
		limit *int,
	) ([]*client.JiraIssue, error)
	GetIssue(
		x context.Context,
		key string,
		comments bool,
	) (*client.JiraIssue, error)
	GetTransitions(
		x context.Context,
		key string,
	) ([]*client.JiraTransition, error)
	TransitionIssue(
		x context.Context,
		key string,
		transitionIdentifier string,
	) error
	AddIssueComment(
		x context.Context,
		key string,
		body string,
	) error
	ListProjects(x context.Context) ([]*client.JiraProject, error)
	CreateIssue(
		x context.Context,
		project string,
		issueType string,
		summary string,
		description string,
		assignee string,
		labels []string,
		fields map[string]any,
	) (*client.JiraIssue, error)
	UpdateIssue(
		x context.Context,
		key string,
		summary string,
		description string,
		assignee string,
		reporter string,
		labels []string,
		fields map[string]any,
		noDiff bool,
	) (*client.IssueUpdateResult, error)
	GetCreateMeta(
		x context.Context,
		project string,
		issueType string,
		expand string,
	) ([]*client.CreateMetaField, error)
	SearchUsers(
		x context.Context,
		query string,
	) ([]*client.JiraUser, error)
	LinkIssues(
		x context.Context,
		key string,
		target string,
		linkType string,
	) error
	DeleteLink(
		x context.Context,
		identifier string,
	) error
	GetLinkTypes(x context.Context) ([]client.LinkType, error)
	UpdateComment(
		x context.Context,
		key string,
		identifier string,
		body string,
	) error
	DeleteComment(
		x context.Context,
		key string,
		identifier string,
	) error
}
