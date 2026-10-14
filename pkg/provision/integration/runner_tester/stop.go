package runner_tester

func (o *Tester) Stop() {
	o.Runner.Close()
}
