package job

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/github/check/job/option"
	"github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/github/run"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
)

func printNotation(
	v []*run.Run,
	o *option.Job,
) {
	r := report.New()

	for _, e := range report.Trim(v, r, o.All, monitorConstant.GoGitHubJob) {
		var s soil.Severity

		if e.HasConcerns() {
			s = soil.Critical
		} else {
			s = soil.Information
		}

		r.AddItem(
			monitorConstant.GoGitHubJob,
			e.MonitorIdentifier,
			s,
			e.Format(constant.NotationFormat),
			*e.Raw.HTMLURL,
			&e.Update,
		)
	}

	r.Print()
}
