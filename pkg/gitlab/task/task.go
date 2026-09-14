package task

import "gitlab.com/gitlab-org/api/client-go/v3"

type Task struct {
	Body  string
	State string
	Link  string
	Type  gitlab.TodoTargetType
	Raw   *gitlab.Todo
}
