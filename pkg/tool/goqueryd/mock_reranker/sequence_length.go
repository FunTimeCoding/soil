package mock_reranker

func (r *Reranker) SequenceLength() int {
	return r.allowance
}
