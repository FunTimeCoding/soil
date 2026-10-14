package reaper

func (r *Reaper) Close() {
	close(r.stop)
}
