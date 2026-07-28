package job

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/github/check/job/option"
	github "github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/github/run"
	monitor "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
)

func printNotation(
	v []*run.Run,
	o *option.Job,
) {
	r := report.New()

	for _, e := range report.Trim(
		v,
		r,
		o.All,
		monitor.GoGitHubJob,
	) {
		var s constant.Severity

		if e.HasConcerns() {
			s = constant.Critical
		} else {
			s = constant.Information
		}

		r.AddItem(
			monitor.GoGitHubJob,
			e.MonitorIdentifier,
			s,
			e.Format(github.NotationFormat),
			*e.Raw.HTMLURL,
			&e.Update,
		)
	}

	r.Print()
}
