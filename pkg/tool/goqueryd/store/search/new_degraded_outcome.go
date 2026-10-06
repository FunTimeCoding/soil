package search

func NewDegradedOutcome(cause error) *Outcome {
	return &Outcome{Degraded: true, Cause: cause}
}
