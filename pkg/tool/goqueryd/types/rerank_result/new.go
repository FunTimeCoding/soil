package rerank_result

func New(
	index int,
	score float64,
) *Result {
	return &Result{Index: index, Score: score}
}
