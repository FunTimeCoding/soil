package service_tester

import "github.com/funtimecoding/soil/pkg/face"

func WithEmbedder(e face.ReachableEmbedder) Option {
	return func(o *Tester) {
		o.embedder = e
	}
}
