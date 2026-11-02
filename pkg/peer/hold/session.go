package hold

import "github.com/funtimecoding/soil/pkg/peer/constant"

func (h *Hold) session() bool {
	e := h.connection.Open(h.context)
	h.reacher.Handle(h.name, e, h.edge, h.refusal)

	if e != nil {
		return false
	}

	defer h.connection.Close()

	if h.context.Err() != nil {
		return true
	}

	reason := constant.ClosedReason

	if f := h.serve(); f != nil {
		reason = f.Error()
	}

	if h.context.Err() != nil {
		return true
	}

	if x := h.reacher.Fail(h.name, reason); x != nil {
		h.edge(x)
	}

	return true
}
