package stale

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/check/stale/option"
	goModule "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/module"
	monitor "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
	"time"
)

func printNotation(
	v []*module.Module,
	o *option.Stale,
) {
	r := report.New()
	f := goModule.Format

	for _, e := range report.Trim(v, r, o.All, monitor.GoStale) {
		s := constant.Information

		if e.HasConcerns() {
			s = constant.Warning
		}

		create := &time.Time{}

		if e.Release != nil {
			create = e.Release
		}

		r.AddItem(
			monitor.GoStale,
			e.MonitorIdentifier,
			s,
			e.Format(f),
			"",
			create,
		)
	}

	r.Print()
}
