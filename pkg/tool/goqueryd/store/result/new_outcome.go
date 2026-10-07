package result

func NewOutcome(results []Search) *Outcome {
	return &Outcome{Results: results}
}
