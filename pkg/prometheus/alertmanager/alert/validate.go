package alert

import (
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"slices"
)

func (a *Alert) Validate() {
	if a.Suppressed() && !slices.Contains(a.concern, constant.Silent) {
		a.concern = append(a.concern, constant.Silent)
	}
}
