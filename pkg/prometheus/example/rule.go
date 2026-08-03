package example

import (
	"fmt"
	console "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"slices"
)

func Rule() {
	f := console.ExtendedColorFormat.Copy()
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
