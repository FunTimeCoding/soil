package reaper

func (r *Reaper) Remove(pid int) {
	r.owned.Delete(pid)
}
