package policy_tester

func (o *Tester) Allows(command string) bool {
	o.t.Helper()

	return o.walk().Allows(command)
}
