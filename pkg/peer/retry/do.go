package retry

import (
	"fmt"
	"time"
)

func Do(
	attempts int,
	backoff time.Duration,
	f func() error,
) (*Result, error) {
	start := time.Now()
	var last error

	for i := range attempts {
		if i > 0 {
			time.Sleep(backoff * time.Duration(1<<(i-1)))
		}

		if last = f(); last == nil {
			return &Result{Attempts: i + 1, Elapsed: time.Since(start)}, nil
		}
	}

	return &Result{Attempts: attempts, Elapsed: time.Since(start)}, fmt.Errorf(
		"after %d attempts: %w",
		attempts,
		last,
	)
}
