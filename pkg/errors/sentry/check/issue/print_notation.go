package issue

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/check/issue/option"
	"github.com/funtimecoding/soil/pkg/errors/sentry/issue"
	monitor "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
)

func printNotation(
	v []*issue.Issue,
	o *option.Issue,
) {
	r := report.New()

	for _, e := range report.Trim(v, r, o.All, monitor.GoSentry) {
		r.AddItem(
			monitor.GoSentry,
			e.MonitorIdentifier,
			constant.Critical,
			e.Title,
			e.Link,
			e.Create,
		)
	}

	r.Print()
}
