package shared_reranker

func (r *Reranker) Allowance() int {
	return r.reranker.Allowance()
}
