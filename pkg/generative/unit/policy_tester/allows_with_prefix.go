package policy_tester

func (o *Tester) AllowsWithPrefix(
	prefix string,
	command string,
) bool {
	o.t.Helper()

	return o.walk().WithScratchPrefixes(prefix).Allows(command)
}
