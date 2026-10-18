package unit

import "context"

func (f *FixedEmbedder) Reachable(context.Context) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return f.failure
}
