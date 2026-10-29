package unit

func (f *FixedEmbedder) FailEmbedding(e error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.embedding = e
}
