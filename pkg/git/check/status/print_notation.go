package status

import (
	"github.com/funtimecoding/soil/pkg/git/check/status/option"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/git/repository"
	monitor "github.com/funtimecoding/soil/pkg/monitor/constant"
	item "github.com/funtimecoding/soil/pkg/monitor/item/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
	"time"
)

func printNotation(
	v []*repository.Repository,
	o *option.Status,
) {
	r := report.New()
	f := constant.Format

	for _, e := range report.Trim(
		v,
		r,
		o.All,
		item.GoGitStatus,
	) {
		var s monitor.Severity

		if e.HasConcerns() {
			s = monitor.Warning
		} else {
			s = monitor.Information
		}

		r.AddItem(
			item.GoGitStatus,
			e.MonitorIdentifier,
			s,
			e.Format(f),
			"",
			&time.Time{},
		)
	}

	r.Print()
}
