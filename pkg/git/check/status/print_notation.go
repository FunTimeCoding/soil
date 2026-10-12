package status

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/git/check/status/option"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/git/repository"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
	"time"
)

func printNotation(
	v []*repository.Repository,
	o *option.Status,
) {
	r := report.New()
	f := constant.Format

	for _, e := range report.Trim(v, r, o.All, monitorConstant.GoGitStatus) {
		var s soil.Severity

		if e.HasConcerns() {
			s = soil.Warning
		} else {
			s = soil.Information
		}

		r.AddItem(
			monitorConstant.GoGitStatus,
			e.MonitorIdentifier,
			s,
			e.Format(f),
			"",
			&time.Time{},
		)
	}

	r.Print()
}
