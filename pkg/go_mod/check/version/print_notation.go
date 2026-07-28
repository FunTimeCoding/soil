package version

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/go_mod"
	"github.com/funtimecoding/soil/pkg/go_mod/check/version/option"
	"github.com/funtimecoding/soil/pkg/go_mod/project"
	monitor "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
	"time"
)

func printNotation(
	v []*project.Project,
	o *option.Version,
) {
	r := report.New()
	f := go_mod.Format

	for _, e := range report.Trim(
		v,
		r,
		o.All,
		monitor.GoGitStatus,
	) {
		var s constant.Severity

		if e.HasConcerns() {
			s = constant.Warning
		} else {
			s = constant.Information
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
