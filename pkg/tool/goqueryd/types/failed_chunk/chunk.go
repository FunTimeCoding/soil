package failed_chunk

type Chunk struct {
	Collection string
	Path       string
	Index      int
	FirstLine  int
	LastLine   int
	Bytes      int
	Reason     string
}
