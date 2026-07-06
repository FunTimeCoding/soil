package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"slices"
)

func Rule() {
	f := option.ExtendedColor.Copy()
	var severities []string
	fmt.Println("Rules")

	for _, r := range prometheus.NewEnvironment().MustRules().Alert() {
		if r.RawAlert != nil {
			fmt.Printf("Alert: %s\n", r.Format(f))

			for k, v := range r.RawAlert.Labels {
				if k == constant.SeverityLabel {
					s := string(v)

					if !slices.Contains(severities, s) {
						severities = append(severities, s)
					}
				}
			}
		} else if r.RawRecord != nil {
			fmt.Printf("Record: %s\n", r.Format(f))

			for k, v := range r.RawRecord.Labels {
				if k == constant.SeverityLabel {
					s := string(v)

					if !slices.Contains(severities, s) {
						severities = append(severities, s)
					}
				}
			}
		} else {
			fmt.Printf("Rule: %+v\n", r)
		}
	}

	fmt.Printf("Severities: %+v\n", severities)
}
