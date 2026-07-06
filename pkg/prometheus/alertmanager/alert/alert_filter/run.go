package alert_filter

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"

func (f *Filter) Run(v []*alert.Alert) []*alert.Alert {
	result := make([]*alert.Alert, 0)

	for _, a := range v {
		if f.match(a) {
			result = append(result, a)
		}
	}

	return result
}
