package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/issue"
	"github.com/funtimecoding/soil/pkg/gitlab/task"
)

type IssueForge interface {
	Issues(x context.Context) ([]*issue.Issue, error)
	MustIssues(x context.Context) []*issue.Issue
	ProjectIssues(
		x context.Context,
		project int64,
	) ([]*issue.Issue, error)
	ProjectsIssues(x context.Context) ([]*issue.Issue, error)
	Tasks(x context.Context) ([]*task.Task, error)
}
