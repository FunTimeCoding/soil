package unit

func (f *FixedEmbedder) Embed(v []string) ([][]float32, error) {
	result := make([][]float32, len(v))

	for i := range v {
		result[i] = []float32{1, 0}
	}

	return result, nil
}
