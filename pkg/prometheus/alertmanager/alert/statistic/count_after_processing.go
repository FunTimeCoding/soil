package statistic

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"strings"
)

func (s *Statistic) CountAfterProcessing(v []*alert.Alert) *Statistic {
	s.Relevant = len(v)

	for _, a := range v {
		s.Group.All++

		switch a.Severity {
		case constant.CriticalSeverity:
			s.Severity.Critical++
		case constant.WarningSeverity:
			s.Severity.Warning++
		case constant.InformationSeverity:
			s.Severity.Information++
		case constant.NoneSeverity:
			s.Severity.None++
		default:
			s.Severity.Unknown++
		}

		switch a.State {
		case constant.ActiveState:
			s.State.Active++
		case constant.SuppressedState:
			s.State.Suppressed++
		}

		if strings.HasPrefix(a.Name, constant.KubernetesPrefix) {
			s.Group.Kubernetes++
		} else {
			s.Group.Other++
		}
	}

	return s
}
