package hold

func (h *Hold) Close() {
	h.cancel()
	h.connection.Close()
	h.started.Do(func() { close(h.done) })
	<-h.done
}
