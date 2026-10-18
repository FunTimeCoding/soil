package retry

import (
	"fmt"
	"time"
)

func Until(
	timeout time.Duration,
	interval time.Duration,
	probe func() error,
) (*Result, error) {
	start := time.Now()
	attempts := 0

	for {
		attempts++
		last := probe()

		if last == nil {
			return &Result{Attempts: attempts, Elapsed: time.Since(start)}, nil
		}

		if time.Since(start)+interval > timeout {
			return &Result{Attempts: attempts, Elapsed: time.Since(start)}, fmt.Errorf(
				"after %s: %w",
				timeout,
				last,
			)
		}

		time.Sleep(interval)
	}
}
