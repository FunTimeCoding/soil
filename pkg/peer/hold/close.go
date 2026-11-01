package hold

func (h *Hold) Close() {
	h.cancel()
	h.connection.Close()
}
