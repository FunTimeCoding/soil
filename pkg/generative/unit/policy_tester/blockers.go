package policy_tester

func (o *Tester) Blockers(command string) []string {
	o.t.Helper()
	w := o.walk()
	w.Allows(command)

	return w.Blockers()
}
