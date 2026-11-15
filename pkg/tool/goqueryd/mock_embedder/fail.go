package mock_embedder

func (f *Embedder) Fail(e error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.failure = e
}
