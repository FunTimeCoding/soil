package job

import (
	"github.com/funtimecoding/soil/pkg/github/check/job/option"
	"github.com/funtimecoding/soil/pkg/github/run"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	item "github.com/funtimecoding/soil/pkg/monitor/item/constant"
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
		item.GoGitHubJob,
	) {
		var s constant.Severity

		if e.HasConcerns() {
			s = constant.Critical
		} else {
			s = constant.Information
		}

		r.AddItem(
			item.GoGitHubJob,
			e.MonitorIdentifier,
			s,
			e.Format(Notation),
			*e.Raw.HTMLURL,
			&e.Update,
		)
	}

	r.Print()
}
