package name_filter

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"

func (f *Filter) Run(v []*alert.Alert) []*alert.Alert {
	var result []*alert.Alert

	for _, a := range v {
		if f.match(a) {
			result = append(result, a)
		}
	}

	return result
}
