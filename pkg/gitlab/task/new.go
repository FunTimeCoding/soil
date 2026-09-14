package task

import "gitlab.com/gitlab-org/api/client-go/v3"

func New(v *gitlab.Todo) *Task {
	return &Task{
		Type:  v.TargetType,
		Body:  v.Body,
		State: v.State,
		Link:  v.TargetURL,
		Raw:   v,
	}
}
