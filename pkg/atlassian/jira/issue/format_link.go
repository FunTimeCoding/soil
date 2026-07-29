package issue

import "github.com/funtimecoding/soil/pkg/atlassian/constant"

func (i *Issue) FormatLink() string {
	if i.Link == "" {
		return constant.JiraNoLink
	}

	return i.Link
}
