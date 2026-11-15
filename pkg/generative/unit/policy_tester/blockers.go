package policy_tester

import "testing"

func Blockers(
	t *testing.T,
	root string,
	scratchpad string,
	command string,
) []string {
	t.Helper()
	w := newWalk(t, root, scratchpad)
	w.Allows(command)

	return w.Blockers()
}
