package result

func NewSearch(
	path string,
	snippet string,
	chunkPosition int,
	snippetLine int,
) *Search {
	return &Search{
		Path:          path,
		Snippet:       snippet,
		ChunkPosition: chunkPosition,
		SnippetLine:   snippetLine,
	}
}
