package runner

import "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"

func (r *Runner) WithInventory(c face.Client) *Runner {
	r.inventory = c

	return r
}
