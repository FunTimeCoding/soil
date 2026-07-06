package alert

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"

func (a *Alert) Suppressed() bool {
	return a.State == constant.SuppressedState
}
