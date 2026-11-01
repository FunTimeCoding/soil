package hold

import "github.com/funtimecoding/soil/pkg/peer/constant"

func (h *Hold) session() bool {
	e := h.connection.Open(h.context)
	h.reacher.Handle(
		h.name,
		e,
		h.logEdge,
		func(f error) { h.reporter.CaptureException(f) },
	)

	if e != nil {
		return false
	}

	defer h.connection.Close()
	reason := constant.ClosedReason

	if f := h.connection.Serve(h.context); f != nil {
		reason = f.Error()
	}

	if h.context.Err() != nil {
		return true
	}

	if x := h.reacher.Fail(h.name, reason); x != nil {
		h.logEdge(x)
	}

	return true
}
