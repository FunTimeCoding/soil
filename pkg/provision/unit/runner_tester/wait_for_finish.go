package runner_tester

import "time"

func (o *Tester) WaitForFinish() {
	o.t.Helper()

	select {
	case <-o.Runner.Finished():
	case <-time.After(5 * time.Second):
		o.t.Fatal("runner did not finish")
	}
}
