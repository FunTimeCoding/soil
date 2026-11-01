package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/classify"
)

func (s *Service) recordGuest(
	name string,
	status string,
) {
	if s.netbox == nil {
		return
	}

	x := context.Background()
	machines, e := s.netboxMachines(x)

	if e == nil && len(machines[name]) > 1 {
		e = ambiguous.Format(
			"netbox holds more than one virtual machine of that name",
		)
	}

	if e == nil {
		e = s.writeGuest(x, name, status)
	}

	if e == nil {
		return
	}

	if classify.Reportable(e) {
		s.reporter.CaptureException(e)

		return
	}

	s.logger.Plain("%s status not recorded: %v", name, e)
}
