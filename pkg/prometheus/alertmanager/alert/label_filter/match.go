package label_filter

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"slices"
)

func (f *Filter) match(a *alert.Alert) bool {
	for k, v := range a.Labels {
		if slices.Contains(f.keep, k) {
			return true
		}

		if slices.Contains(f.drop, k) {
			return false
		}

		for _, l := range f.keepValue {
			if l.Match(k, v) {
				return true
			}
		}

		for _, l := range f.dropValue {
			if l.Match(k, v) {
				return false
			}
		}
	}

	return f.KeepByDefault
}
