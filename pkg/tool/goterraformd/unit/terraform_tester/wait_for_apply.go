package terraform_tester

func (o *Tester) waitForApply() {
	o.t.Helper()
	o.WaitForRuns(1)
}
