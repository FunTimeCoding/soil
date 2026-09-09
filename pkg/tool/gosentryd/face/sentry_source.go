package face

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

type SentrySource interface {
	Organizations() ([]response.Organization, error)
	OrganizationProjects(organization string) ([]response.Project, error)
	Releases(
		organization string,
		query string,
		limit int,
	) ([]response.Release, error)
	SearchIssues(
		organization string,
		query string,
		project string,
		limit int,
		cursor string,
	) ([]response.Issue, error)
	SearchEvents(
		organization string,
		query string,
		project string,
		limit int,
		cursor string,
	) ([]response.EventRow, error)
	IssueByIdentifier(
		organization string,
		identifier string,
	) (*response.Issue, error)
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
	LatestEvent(
		organization string,
		issueIdentifier string,
	) (*response.Event, error)
	Event(
		organization string,
		project string,
		identifier string,
	) (*response.Event, error)
	UpdateIssue(
		organization string,
		identifier string,
		status string,
		assignedTo string,
	) (*response.Issue, error)
	DeleteIssue(
		organization string,
		identifier string,
	) error
	Whoami() (*response.User, error)
}
