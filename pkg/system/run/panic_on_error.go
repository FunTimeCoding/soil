package run

import "github.com/funtimecoding/soil/pkg/errors"

func (r *Run) PanicOnError() {
	errors.PanicOnError(r.Error)
}
