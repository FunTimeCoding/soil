package face

import (
	"github.com/funtimecoding/soil/pkg/github/issue"
	"github.com/funtimecoding/soil/pkg/github/pull_request"
	"github.com/google/go-github/v93/github"
)

type IssueClient interface {
	BranchRequest(
		owner string,
		repository string,
		branch string,
	) (*pull_request.Request, error)
	FindBranchRequest(
		owner string,
		repository string,
		branch string,
	) (*pull_request.Request, bool, error)
	Issues() ([]*issue.Issue, error)
	MustBranchRequest(
		owner string,
		repository string,
		branch string,
	) *pull_request.Request
	MustProjectIssues(
		owner string,
		name string,
	) []*github.Issue
	MustPullRequests(
		owner string,
		name string,
	) []*github.PullRequest
	OpenUserIssues() ([]*issue.Issue, error)
	ProjectIssues(
		owner string,
		name string,
	) ([]*github.Issue, error)
	PullRequests(
		owner string,
		name string,
	) ([]*github.PullRequest, error)
	SearchIssue(
		query string,
		a ...any,
	) ([]*issue.Issue, error)
}
