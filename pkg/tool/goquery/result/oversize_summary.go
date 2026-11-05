package result

type OversizeSummary struct {
	Files        int `json:"files"`
	Chunks       int `json:"chunks"`
	OverEmbedder int `json:"over_embedder"`
}
