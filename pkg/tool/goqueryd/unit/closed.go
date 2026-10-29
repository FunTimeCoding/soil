package unit

func (f *FixedEmbedder) Closed() int {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return f.closed
}
