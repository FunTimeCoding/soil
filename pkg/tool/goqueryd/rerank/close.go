package rerank

import "github.com/funtimecoding/soil/pkg/errors"

func (r *Reranker) Close() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.session != nil {
		destroySession(r.session)
		r.session = nil
	}

	if r.tokenizer != nil {
		errors.PanicOnError(r.tokenizer.Close())
		r.tokenizer = nil
	}

	r.counterMutex.Lock()
	defer r.counterMutex.Unlock()

	if r.counter != nil {
		errors.PanicOnError(r.counter.Close())
		r.counter = nil
	}

	return nil
}
