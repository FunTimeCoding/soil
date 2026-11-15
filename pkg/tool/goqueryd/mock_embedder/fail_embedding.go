package mock_embedder

func (f *Embedder) FailEmbedding(e error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.embedding = e
}
