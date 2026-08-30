package store

type SearchResult struct {
	VirtualPath   string              `json:"virtual_path"`
	FilePath      string              `json:"file_path"`
	Collection    string              `json:"collection"`
	Path          string              `json:"path"`
	Title         string              `json:"title"`
	Hash          string              `json:"hash"`
	Score         float64             `json:"score"`
	Source        string              `json:"source"`
	SourceType    string              `json:"source_type,omitempty"`
	Context       string              `json:"context,omitempty"`
	Snippet       string              `json:"snippet,omitempty"`
	SnippetLine   int                 `json:"snippet_line,omitzero"`
	Body          string              `json:"body,omitempty"`
	Metadata      map[string][]string `json:"metadata,omitempty"`
	ChunkPosition int                 `json:"-"`
}
