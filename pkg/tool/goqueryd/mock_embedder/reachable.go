package mock_embedder

import "context"

func (f *Embedder) Reachable(context.Context) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return f.failure
}
