package rerank

import "github.com/funtimecoding/soil/pkg/errors"

func (r *Reranker) Close() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, s := range r.sessions {
		destroySession(s)
	}

	r.sessions = nil

	if r.tokenizer != nil {
		errors.PanicOnError(r.tokenizer.Close())
		r.tokenizer = nil
	}

	return nil
}
