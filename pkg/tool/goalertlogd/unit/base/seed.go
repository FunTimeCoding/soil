package base

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
)

func (s *Server) Seed(count int) {
	for i := range count {
		s.MockClient.Add(
			alert.NewBasic(
				fmt.Sprintf("fingerprint-%d", i),
				"HypervisorUnreachable",
				constant.WarningSeverity,
				"",
			),
		)
	}

	s.Worker.Poll()
}
