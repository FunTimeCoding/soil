package watcher

import "context"

func (w *Watcher) attempt(call func() error) error {
	return w.guard.Attempt(
		context.Background(),
		func(context.Context) error {
			return call()
		},
	)
}
