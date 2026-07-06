package alert

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"slices"
)

func FilterSevere(v []*Alert) []*Alert {
	var result []*Alert

	for _, e := range v {
		if e.State != constant.ActiveState {
			continue
		}

		if !slices.Contains(constant.SevereSeverities, e.Severity) {
			continue
		}

		result = append(result, e)
	}

	return result
}
