package service_tester

import "github.com/funtimecoding/soil/pkg/face"

func WithNotifier(n face.EventNotifier) Option {
	return func(o *Tester) {
		o.notifier = n
	}
}
