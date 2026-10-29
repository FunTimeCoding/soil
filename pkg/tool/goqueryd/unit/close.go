package unit

func (f *FixedEmbedder) Close() {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.closed++
}
