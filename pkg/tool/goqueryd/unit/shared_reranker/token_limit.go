package shared_reranker

func (r *Reranker) TokenLimit() int {
	return r.reranker.TokenLimit()
}
