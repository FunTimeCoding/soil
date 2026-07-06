package alert

import "github.com/funtimecoding/soil/pkg/prometheus/constant"

func (a *Alert) IsKubernetes() bool {
	return a.Detail(constant.NamespaceLabel) != ""
}
