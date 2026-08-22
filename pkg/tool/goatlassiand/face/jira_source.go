package face

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/basic"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/field_map"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
)

type JiraSource interface {
	Issue(key string) (*issue.Issue, error)
	Project(key string) (*jira.Project, error)
	Projects() (*jira.ProjectList, error)
	FieldMap() (*field_map.Map, error)
	Comment(
		key string,
		body string,
	) error
	Transition(
		key string,
		transitionIdentifier string,
	) error
	Transitions(key string) ([]jira.Transition, error)
	Search(
		query string,
		a ...any,
	) ([]*issue.Issue, error)
	SearchLimit(
		limit int,
		query string,
		a ...any,
	) ([]*issue.Issue, error)
	User() (*jira.User, error)
	MetaProject(key string) (*jira.MetaProject, error)
	MetaIssueType(
		p *jira.MetaProject,
		issueType string,
	) (*jira.MetaIssueType, error)
	CreateNative(i *jira.Issue) (*issue.Issue, error)
	UpdateNative(i *jira.Issue) error
	Assign(
		key string,
		u *jira.User,
	) error
	AddLink(link *jira.IssueLink) error
	DeleteLink(identifier string) error
	UpdateComment(
		key string,
		comment *jira.Comment,
	) error
	DeleteComment(
		key string,
		identifier string,
	) error
	FindUsers(query string) ([]jira.User, error)
	Basic() *basic.Client
}
