package pull_request

import (
	"github.com/funtimecoding/soil/pkg/github/check/pull_request/option"
	"github.com/funtimecoding/soil/pkg/github/run"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	item "github.com/funtimecoding/soil/pkg/monitor/item/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
)

func printNotation(
	v []*run.Run,
	o *option.Request,
) {
	r := report.New()

	for _, e := range report.Trim(
		v,
		r,
		o.All,
		item.GoGitHubPullRequest,
	) {
		var s constant.Severity

		if e.HasConcerns() {
			s = constant.Critical
		} else {
			s = constant.Information
		}

		r.AddItem(
			item.GoGitHubPullRequest,
			e.MonitorIdentifier,
			s,
			e.Format(Notation),
			*e.Raw.HTMLURL,
			&e.Update,
		)
	}

	r.Print()
}
