package failed_chunk

func New(
	collection string,
	path string,
	index int,
	reason string,
) *Chunk {
	return &Chunk{
		Collection: collection,
		Path:       path,
		Index:      index,
		Reason:     reason,
	}
}
