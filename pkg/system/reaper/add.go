package reaper

func (r *Reaper) Add(pid int) {
	r.owned.Store(pid, struct{}{})
}
