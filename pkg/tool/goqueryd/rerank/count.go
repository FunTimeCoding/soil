package rerank

import (
	"github.com/amikos-tech/pure-tokenizers"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (r *Reranker) Count(text string) int {
	r.counterMutex.Lock()
	defer r.counterMutex.Unlock()
	result, e := r.counter.Encode(
		valid(text),
		tokenizers.WithAddSpecialTokens(),
	)
	errors.PanicOnError(e)

	return len(result.IDs) - r.singleSpecials
}
