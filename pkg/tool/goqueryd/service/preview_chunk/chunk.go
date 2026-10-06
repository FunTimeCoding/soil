package preview_chunk

type Chunk struct {
	Index      int
	FirstLine  int
	LastLine   int
	Bytes      int
	Tokens     int
	Piece      bool
	CutChecked bool
	CutLevel   int
	CutLines   []int
}
