package service_tester

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"

func WithEmbedCounter(c face.EmbedCounter) Option {
	return func(o *Tester) {
		o.counter = c
	}
}
