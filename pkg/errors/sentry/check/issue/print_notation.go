package issue

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/check/issue/option"
	"github.com/funtimecoding/soil/pkg/errors/sentry/issue"
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
		item.GoSentry,
	) {
		r.AddItem(
			item.GoSentry,
			e.MonitorIdentifier,
			constant.Critical,
			e.Title,
			e.Link,
			e.Create,
		)
	}

	r.Print()
}
