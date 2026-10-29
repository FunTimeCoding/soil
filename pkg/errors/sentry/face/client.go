package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/errors/sentry/issue"
)

type Client interface {
	AllIssues() ([]*issue.Issue, error)
	Close()
	DeleteIssue(
		organization string,
		identifier string,
	) error
	Event(
		organization string,
		project string,
		identifier string,
	) (*response.Event, error)
	IssueByIdentifier(
		organization string,
		identifier string,
	) (*response.Issue, error)
	IssueByShortIdentifier(
		organization string,
		identifier string,
	) (*issue.Issue, error)
	IssueEvents(
		organization string,
		identifier string,
		query string,
		limit int,
		cursor string,
	) ([]response.Event, error)
	IssueTagValues(
		organization string,
		identifier string,
		tag string,
		limit int,
	) ([]response.TagValue, error)
	Issues(
		organization string,
		projectIdentifier string,
		period string,
	) ([]*issue.Issue, error)
	IssuesSimple(verbose bool) ([]*issue.Issue, error)
	LatestEvent(
		organization string,
		issueIdentifier string,
	) (*response.Event, error)
	MustIssueByShortIdentifier(
		organization string,
		identifier string,
	) *issue.Issue
	MustIssues(
		organization string,
		projectIdentifier string,
		period string,
	) []*issue.Issue
	MustIssuesSimple(verbose bool) []*issue.Issue
	MustLatestEvent(
		organization string,
		issueIdentifier string,
	) *response.Event
	MustOrganizationProjects(organization string) []response.Project
	MustOrganizations() []response.Organization
	MustProjects() []response.Project
	MustTeams(organization string) []response.Team
	MustTrackedIssues() []*issue.Issue
	MustWhoami() *response.User
	Organization(slug string) (*response.Organization, error)
	OrganizationProjects(organization string) ([]response.Project, error)
	Organizations() ([]response.Organization, error)
	Project(
		organization string,
		projectSlug string,
	) (*response.Project, error)
	Projects() ([]response.Project, error)
	Reachable(x context.Context) error
	Releases(
		organization string,
		query string,
		limit int,
	) ([]response.Release, error)
	SearchEvents(
		organization string,
		query string,
		project string,
		limit int,
		cursor string,
	) ([]response.EventRow, error)
	SearchIssues(
		organization string,
		query string,
		project string,
		limit int,
		cursor string,
	) ([]response.Issue, error)
	Teams(organization string) ([]response.Team, error)
	TrackedIssues() ([]*issue.Issue, error)
	UpdateIssue(
		organization string,
		identifier string,
		status string,
		assignedTo string,
	) (*response.Issue, error)
	Whoami() (*response.User, error)
}
