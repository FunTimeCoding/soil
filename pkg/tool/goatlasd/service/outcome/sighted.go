package outcome

import "time"

func Sighted(
	source string,
	count int,
	swept int64,
	took time.Duration,
) *Outcome {
	return &Outcome{
		Source:    source,
		Sightings: true,
		Count:     count,
		Swept:     swept,
		Took:      took,
	}
}
