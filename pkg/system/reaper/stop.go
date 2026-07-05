package reaper

func (r *Reaper) Stop() {
	close(r.stop)
}
