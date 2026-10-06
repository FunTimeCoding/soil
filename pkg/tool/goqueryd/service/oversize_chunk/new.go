package oversize_chunk

func New(
	index int,
	firstLine int,
	lastLine int,
	bytes int,
	tokens int,
) *Chunk {
	return &Chunk{
		Index:     index,
		FirstLine: firstLine,
		LastLine:  lastLine,
		Bytes:     bytes,
		Tokens:    tokens,
	}
}
