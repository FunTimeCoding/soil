package shared_reranker

func (r *Reranker) Count(text string) int {
	return r.reranker.Count(text)
}
