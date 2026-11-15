package mock_embedder

func (f *Embedder) Calls() int {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return f.calls
}
