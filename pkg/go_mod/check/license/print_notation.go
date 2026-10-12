package license

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/check/license/option"
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/dependency"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
	"time"
)

func printNotation(
	v []*dependency.Dependency,
	o *option.License,
) {
	r := report.New()
	f := constant.Format

	for _, e := range report.Trim(v, r, o.All, monitorConstant.GoLicense) {
		var s soil.Severity

		if e.HasConcerns() {
			s = soil.Warning
		} else {
			s = soil.Information
		}

		r.AddItem(
			monitorConstant.GoLicense,
			e.MonitorIdentifier,
			s,
			e.Format(f),
			"",
			&time.Time{},
		)
	}

	r.Print()
}
