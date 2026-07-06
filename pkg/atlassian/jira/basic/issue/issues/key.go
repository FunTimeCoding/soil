package issues

import "github.com/funtimecoding/soil/pkg/atlassian/jira/basic/issue"

func (i *Issues) Key(k string) *issue.Issue {
	for _, s := range i.list {
		if s.Key == k {
			return s
		}
	}

	return nil
}
