package unit

func (f *FixedEmbedder) Refuse(marker string) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.refused = marker
}
