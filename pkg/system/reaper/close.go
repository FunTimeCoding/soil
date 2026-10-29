package reaper

func (r *Reaper) Close() {
	r.closing.Do(func() { close(r.stop) })
}
