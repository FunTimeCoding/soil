package outcome

import "time"

func Placed(
	source string,
	count int,
	swept int64,
	took time.Duration,
) *Outcome {
	return &Outcome{Source: source, Count: count, Swept: swept, Took: took}
}
