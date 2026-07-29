package issue

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"slices"
)

func (i *Issue) Blocked() bool {
	for _, l := range i.Raw.Fields.IssueLinks {
		if l.InwardIssue == nil {
			continue
		}

		if l.Type.Inward == constant.JiraBlockedBy {
			if slices.Contains(
				i.option.ClosedStatus,
				l.InwardIssue.Fields.Status.Name,
			) {
				continue
			}

			return true
		}
	}

	return false
}
