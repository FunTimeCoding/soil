package search

func NewResult(
	path string,
	snippet string,
	chunkPosition int,
	snippetLine int,
) *Result {
	return &Result{
		Path:          path,
		Snippet:       snippet,
		ChunkPosition: chunkPosition,
		SnippetLine:   snippetLine,
	}
}
