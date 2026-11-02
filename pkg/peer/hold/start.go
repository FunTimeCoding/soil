package hold

func (h *Hold) Start() {
	h.started.Do(func() { go h.loop() })
}
