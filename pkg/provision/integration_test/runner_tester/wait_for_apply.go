package runner_tester

import "time"

func (o *Tester) WaitForApply(count int) {
	o.t.Helper()
	deadline := time.Now().Add(5 * time.Second)

	for time.Now().Before(deadline) {
		if o.ApplyCount() >= count {
			return
		}

		time.Sleep(10 * time.Millisecond)
	}

	o.t.Fatalf(
		"timed out waiting for %d applies, got %d",
		count,
		o.ApplyCount(),
	)
}
