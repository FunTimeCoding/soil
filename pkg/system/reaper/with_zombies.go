package reaper

import "github.com/funtimecoding/soil/pkg/system/types/zombie_detail"

func (r *Reaper) WithZombies(f func() map[int]zombie_detail.Detail) *Reaper {
	r.zombies = f

	return r
}
