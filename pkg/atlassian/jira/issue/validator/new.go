package validator

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func New(_ *issue.Issue) *Validator {
	return &Validator{}
}
