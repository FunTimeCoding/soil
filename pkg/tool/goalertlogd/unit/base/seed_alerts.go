package base

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	prometheusConstant "github.com/funtimecoding/soil/pkg/prometheus/constant"
)

func (s *Server) SeedAlerts() {
	s.t.Helper()
	s.Alertmanager.Add(
		alert.NewBasic(
			"fp1",
			"HighMemory",
			prometheusConstant.CriticalSeverity,
			"Memory above 90%",
		),
	)
	s.Alertmanager.Add(
		alert.NewBasic(
			"fp2",
			"DiskFull",
			prometheusConstant.WarningSeverity,
			"Disk usage above 85%",
		),
	)
	s.Clock().Poll()
}
