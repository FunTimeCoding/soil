package issue

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/check/issue/option"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	soil "github.com/funtimecoding/soil/pkg/constant"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
)

func printNotation(
	v []*issue.Issue,
	o *option.Issue,
) {
	r := report.New()

	for _, e := range report.Trim(v, r, o.All, monitorConstant.GoJira) {
		r.AddItem(
			monitorConstant.GoJira,
			e.MonitorIdentifier,
			soil.Warning,
			e.Summary,
			e.Link,
			e.Create,
		)
	}

	r.Print()
}
