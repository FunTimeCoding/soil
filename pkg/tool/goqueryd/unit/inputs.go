package unit

func (f *FixedEmbedder) Inputs() []string {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return append([]string{}, f.inputs...)
}
