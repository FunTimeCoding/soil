package status

type Status struct {
	Collections       []CollectionStatus `json:"collections"`
	TotalDocuments    int                `json:"total_documents"`
	TotalEmbeddings   int                `json:"total_embeddings"`
	PendingEmbeddings int                `json:"pending_embeddings"`
	FailedEmbeddings  int                `json:"failed_embeddings"`
}
