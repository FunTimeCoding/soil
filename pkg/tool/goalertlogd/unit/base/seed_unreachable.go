package base

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	prometheusConstant "github.com/funtimecoding/soil/pkg/prometheus/constant"
)

func (s *Server) SeedUnreachable(count int) {
	for i := range count {
		s.Alertmanager.Add(
			alert.NewBasic(
				fmt.Sprintf("fingerprint-%d", i),
				"HypervisorUnreachable",
				prometheusConstant.WarningSeverity,
				"",
			),
		)
	}

	s.Clock().Poll()
}
