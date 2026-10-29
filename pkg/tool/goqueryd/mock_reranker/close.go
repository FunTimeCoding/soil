package mock_reranker

func (r *Reranker) Close() error {
	r.Closed++

	return nil
}
