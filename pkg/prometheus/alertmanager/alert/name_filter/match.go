package name_filter

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"slices"
)

func (f *Filter) match(a *alert.Alert) bool {
	if slices.Contains(f.keep, a.Name) {
		return true
	}

	if slices.Contains(f.drop, a.Name) {
		return false
	}

	return f.KeepByDefault
}
