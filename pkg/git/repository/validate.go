package repository

import "github.com/funtimecoding/soil/pkg/git/constant"

func (r *Repository) Validate() {
	if !r.IsClean {
		r.concern = append(r.concern, constant.NotClean)
	}
}
