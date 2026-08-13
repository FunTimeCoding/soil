package status

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/git/check/status/option"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/git/repository"
	monitor "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
	"time"
)

func printNotation(
	v []*repository.Repository,
	o *option.Status,
) {
	r := report.New()
	f := constant.Format

	for _, e := range report.Trim(v, r, o.All, monitor.GoGitStatus) {
		var s library.Severity

		if e.HasConcerns() {
			s = library.Warning
		} else {
			s = library.Information
		}

		r.AddItem(
			monitor.GoGitStatus,
			e.MonitorIdentifier,
			s,
			e.Format(f),
			"",
			&time.Time{},
		)
	}

	r.Print()
}
