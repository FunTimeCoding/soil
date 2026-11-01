package hold

import "github.com/funtimecoding/soil/pkg/peer/reacher/edge"

func (h *Hold) WithReport(
	crossed func(*edge.Edge),
	refused func(error),
) *Hold {
	h.edge = crossed
	h.refusal = refused

	return h
}
