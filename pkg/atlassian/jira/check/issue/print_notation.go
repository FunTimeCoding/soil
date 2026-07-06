package issue

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/check/issue/option"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	item "github.com/funtimecoding/soil/pkg/monitor/item/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
)

func printNotation(
	v []*issue.Issue,
	o *option.Issue,
) {
	r := report.New()

	for _, e := range report.Trim(
		v,
		r,
		o.All,
		item.GoJira,
	) {
		r.AddItem(
			item.GoJira,
			e.MonitorIdentifier,
			constant.Warning,
			e.Summary,
			e.Link,
			e.Create,
		)
	}

	r.Print()
}
