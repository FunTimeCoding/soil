package shared_reranker

func Lease() *Reranker {
	return &Reranker{reranker: shared()}
}
