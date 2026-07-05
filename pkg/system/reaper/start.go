package reaper

func (r *Reaper) Start() {
	go r.loop()
}
