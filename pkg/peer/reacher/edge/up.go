package edge

import "time"

func Up(
	host string,
	duration time.Duration,
) *Edge {
	return &Edge{Host: host, Duration: duration}
}
