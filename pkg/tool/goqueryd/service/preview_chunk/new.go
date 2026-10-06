package preview_chunk

func New(
	index int,
	firstLine int,
	lastLine int,
	bytes int,
	tokens int,
	piece bool,
) *Chunk {
	return &Chunk{
		Index:     index,
		FirstLine: firstLine,
		LastLine:  lastLine,
		Bytes:     bytes,
		Tokens:    tokens,
		Piece:     piece,
	}
}
