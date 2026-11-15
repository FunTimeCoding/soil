package mock_embedder

func (f *Embedder) Close() {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.closed++
}
