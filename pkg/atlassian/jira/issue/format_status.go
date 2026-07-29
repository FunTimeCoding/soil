package issue

import "github.com/funtimecoding/soil/pkg/atlassian/constant"

func (i *Issue) FormatStatus() string {
	if i.Status == "" {
		return constant.JiraNoStatus
	}

	if i.shortStatus != nil {
		return i.shortStatus(i.Status)
	}

	return i.Status
}
