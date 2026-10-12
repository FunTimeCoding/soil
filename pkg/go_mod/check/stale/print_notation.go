package stale

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/check/stale/option"
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/module"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
	"time"
)

func printNotation(
	v []*module.Module,
	o *option.Stale,
) {
	r := report.New()
	f := constant.Format

	for _, e := range report.Trim(v, r, o.All, monitorConstant.GoStale) {
		s := soil.Information

		if e.HasConcerns() {
			s = soil.Warning
		}

		create := &time.Time{}

		if e.Release != nil {
			create = e.Release
		}

		r.AddItem(
			monitorConstant.GoStale,
			e.MonitorIdentifier,
			s,
			e.Format(f),
			"",
			create,
		)
	}

	r.Print()
}
