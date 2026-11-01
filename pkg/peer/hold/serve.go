package hold

func (h *Hold) serve() error {
	h.mutex.Lock()
	h.serving = true
	h.mutex.Unlock()
	defer func() {
		h.mutex.Lock()
		h.serving = false
		h.mutex.Unlock()
	}()

	return h.connection.Serve(h.context)
}
