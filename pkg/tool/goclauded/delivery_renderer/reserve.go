package delivery_renderer

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (r *Renderer) reserve(entries []queue.Entry) int {
	var result int

	for _, e := range entries {
		result += r.smallest(e)
	}

	return result
}
