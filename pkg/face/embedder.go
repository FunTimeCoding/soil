package face

type Embedder interface {
	Embed(v []string) ([][]float32, error)
}
