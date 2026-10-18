package retry

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func Until(
	x context.Context,
	timeout time.Duration,
	interval time.Duration,
	probe func(context.Context) error,
) (*Result, error) {
	start := time.Now()
	attempts := 0
	bounded, cancel := context.WithTimeout(x, timeout)
	defer cancel()

	for {
		attempts++
		last := probe(bounded)
		result := &Result{Attempts: attempts, Elapsed: time.Since(start)}

		if last == nil {
			return result, nil
		}

		select {
		case <-bounded.Done():
			if errors.Is(bounded.Err(), context.DeadlineExceeded) {
				return result, fmt.Errorf("after %s: %w", timeout, last)
			}

			return result, fmt.Errorf("%w: %w", bounded.Err(), last)
		case <-time.After(interval):
		}
	}
}
