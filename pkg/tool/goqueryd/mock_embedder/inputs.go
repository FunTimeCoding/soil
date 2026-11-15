package mock_embedder

func (f *Embedder) Inputs() []string {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return append([]string{}, f.inputs...)
}
