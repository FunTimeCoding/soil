package issues

import "github.com/funtimecoding/soil/pkg/atlassian/jira/basic/issue"

func (i *Issues) Add(v *issue.Issue) {
	i.list = append(i.list, v)
}
