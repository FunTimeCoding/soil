package shared_reranker

import "github.com/funtimecoding/soil/pkg/errors"

func Release() {
	if loaded.Load() {
		errors.PanicClose(shared())
	}
}
