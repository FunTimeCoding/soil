package runner

import "context"

func (r *Runner) attempt(call func() error) error {
	return r.guard.Attempt(
		context.Background(),
		func(context.Context) error {
			return call()
		},
	)
}
