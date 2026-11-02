package alert

import prometheusConstant "github.com/funtimecoding/soil/pkg/prometheus/constant"

func (a *Alert) IsKubernetes() bool {
	return a.Detail(prometheusConstant.NamespaceLabel) != ""
}
