package face

import (
	"github.com/funtimecoding/soil/pkg/gitlab/issue"
	"github.com/funtimecoding/soil/pkg/gitlab/task"
)

type IssueForge interface {
	Issues() ([]*issue.Issue, error)
	MustIssues() []*issue.Issue
	ProjectIssues(project int64) ([]*issue.Issue, error)
	ProjectsIssues() ([]*issue.Issue, error)
	Tasks() ([]*task.Task, error)
}
