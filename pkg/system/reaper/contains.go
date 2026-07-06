package reaper

func (r *Reaper) Contains(pid int) bool {
	_, okay := r.owned.Load(pid)

	return okay
}
