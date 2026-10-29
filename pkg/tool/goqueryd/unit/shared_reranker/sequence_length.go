package shared_reranker

func (r *Reranker) SequenceLength() int {
	return r.reranker.SequenceLength()
}
