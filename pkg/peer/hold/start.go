package hold

func (h *Hold) Start() {
	go h.loop()
}
