package runner_tester

func (o *Tester) ExpectFinished() {
	o.t.Helper()

	select {
	case <-o.Runner.Finished():
	default:
		o.t.Fatal("runner still running")
	}
}
