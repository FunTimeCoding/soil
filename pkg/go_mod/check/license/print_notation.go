package license

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/check/license/option"
	go_modConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/dependency"
	monitor "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
	"time"
)

func printNotation(
	v []*dependency.Dependency,
	o *option.License,
) {
	r := report.New()
	f := go_modConstant.Format

	for _, e := range report.Trim(v, r, o.All, monitor.GoLicense) {
		var s constant.Severity

		if e.HasConcerns() {
			s = constant.Warning
		} else {
			s = constant.Information
		}

		r.AddItem(
			monitor.GoLicense,
			e.MonitorIdentifier,
			s,
			e.Format(f),
			"",
			&time.Time{},
		)
	}

	r.Print()
}
