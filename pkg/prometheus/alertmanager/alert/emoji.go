package alert

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"

func (a *Alert) emoji() []string {
	var result []string

	if a.Suppressed() {
		result = append(result, "🔇")
	}

	if a.Runbook != constant.None {
		result = append(result, "🖊 ")
	}

	return result
}
