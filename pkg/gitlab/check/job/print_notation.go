package job

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/check/job/option"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	monitor "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
)

func printNotation(
	v []*job.Job,
	o *option.Job,
) {
	r := report.New()
	f := constant.CheckFormat

	for _, e := range report.Trim(
		v,
		r,
		o.All,
		monitor.GoGitLab,
	) {
		var s library.Severity

		if e.Fail() {
			s = library.Critical
		} else {
			s = library.Information
		}

		r.AddItem(
			monitor.GoGitLab,
			e.MonitorIdentifier,
			s,
			e.Format(f),
			e.Link,
			e.Create,
		)
	}

	r.Print()
}
