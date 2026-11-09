package run

import "context"

func (r *Run) WithContext(x context.Context) *Run {
	r.context = x

	return r
}
