package record

type EmbeddingCandidate struct {
	FilePath   string
	Collection string
	Path       string
	Title      string
	Hash       string
	Body       string
	Position   int
	Vector     []float32
	Distance   float64
}
