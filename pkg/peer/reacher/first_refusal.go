package reacher

import "github.com/funtimecoding/soil/pkg/strings/join"

func (r *Reacher) firstRefusal(
	name string,
	e error,
) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	key := join.Colon(name, e.Error())

	if r.refused[key] {
		return false
	}

	r.refused[key] = true

	return true
}
