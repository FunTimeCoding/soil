package mock_embedder

func (f *Embedder) Refuse(marker string) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.refused = marker
}
