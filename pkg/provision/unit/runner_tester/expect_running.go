package runner_tester

import "time"

func (o *Tester) ExpectRunning() {
	o.t.Helper()

	select {
	case <-o.Runner.Finished():
		o.t.Fatal("runner finished with an apply in flight")
	case <-time.After(100 * time.Millisecond):
	}
}
