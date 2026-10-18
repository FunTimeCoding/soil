package unit

func (f *FixedEmbedder) Fail(e error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.failure = e
}
