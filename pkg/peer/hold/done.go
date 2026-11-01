package hold

func (h *Hold) Done() <-chan struct{} {
	return h.done
}
