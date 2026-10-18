package unit

func (f *FixedEmbedder) Calls() int {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return f.calls
}
