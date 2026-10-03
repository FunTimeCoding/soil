package alert

import "k8s.io/apimachinery/pkg/util/duration"

func (a *Alert) formatAge() string {
	return duration.HumanDuration(a.Age())
}
