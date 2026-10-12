package version

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/check/version/option"
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/project"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
	"time"
)

func printNotation(
	v []*project.Project,
	o *option.Version,
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
