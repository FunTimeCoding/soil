package service_tester

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"

func WithReranker(r face.Reranker) Option {
	return func(o *Tester) {
		o.reranker = r
	}
}
