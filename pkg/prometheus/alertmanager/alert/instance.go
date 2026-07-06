package alert

import "github.com/funtimecoding/soil/pkg/prometheus/constant"

func (a *Alert) Instance() string {
	result := a.Detail(constant.InstanceLabel)

	if result != "" && a.instance != nil {
		return a.instance(result)
	}

	return result
}
