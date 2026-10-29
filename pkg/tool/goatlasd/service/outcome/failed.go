package outcome

import "time"

func Failed(
	source string,
	took time.Duration,
	e error,
) *Outcome {
	return &Outcome{Source: source, Took: took, Error: e}
}
