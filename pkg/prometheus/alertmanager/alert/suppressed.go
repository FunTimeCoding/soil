package alert

import "github.com/funtimecoding/soil/pkg/prometheus/constant"

func (a *Alert) Suppressed() bool {
	return a.State == constant.SuppressedState
}
