package runner

import "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"

func (r *Runner) WithInventory(c face.InventoryMarkup) *Runner {
	r.inventory = c

	return r
}
