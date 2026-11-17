package hold

import "github.com/funtimecoding/soil/pkg/peer/outage_tracker/edge"

func (h *Hold) WithReport(
	crossed func(*edge.Edge),
	refused func(error),
) *Hold {
	h.edge = crossed
	h.refusal = refused

	return h
}
