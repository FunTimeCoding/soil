package mock_embedder

func (f *Embedder) Closed() int {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return f.closed
}
