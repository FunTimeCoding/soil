package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"slices"
)

func Rule() {
	f := consoleConstant.ExtendedColorFormat.Copy()
	var severities []string
	console.Line("Rules")

	for _, r := range prometheus.NewEnvironment().MustRules().Alert() {
		if r.RawAlert != nil {
			console.Format("Alert: %s\n", r.Format(f))

			for k, v := range r.RawAlert.Labels {
				if k == constant.SeverityLabel {
					s := string(v)

					if !slices.Contains(severities, s) {
						severities = append(severities, s)
					}
				}
			}
		} else if r.RawRecord != nil {
			console.Format("Record: %s\n", r.Format(f))

			for k, v := range r.RawRecord.Labels {
				if k == constant.SeverityLabel {
					s := string(v)

					if !slices.Contains(severities, s) {
						severities = append(severities, s)
					}
				}
			}
		} else {
			console.Format("Rule: %+v\n", r)
		}
	}

	console.Format("Severities: %+v\n", severities)
}
