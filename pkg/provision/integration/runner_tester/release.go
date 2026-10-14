package runner_tester

func (o *Tester) Release() {
	o.release.Do(
		func() {
			close(o.gate)
		},
	)
}
