package delivery

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (d *Delivery) reserve(entries []queue.Entry) int {
	var result int

	for _, e := range entries {
		result += d.smallest(e)
	}

	return result
}
