package search

func NewOutcome(results []Result) *Outcome {
	return &Outcome{Results: results}
}
