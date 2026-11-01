package plate_page

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
)

func statusColor(i *issue.Issue) string {
	if i.Blocked() {
		return webConstant.BadColor
	}

	switch i.Raw.Fields.Status.StatusCategory.Key {
	case "indeterminate":
		return webConstant.BusyColor
	case "new":
		return webConstant.WarnColor
	}

	return webConstant.QuietColor
}
