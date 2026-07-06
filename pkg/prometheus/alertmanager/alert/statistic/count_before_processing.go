package statistic

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"

func (s *Statistic) CountBeforeProcessing(v []*alert.Alert) *Statistic {
	s.Total = len(v)

	return s
}
