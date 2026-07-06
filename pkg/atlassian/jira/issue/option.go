package issue

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue/option"

func (i *Issue) Option() *option.Issue {
	return i.option
}
