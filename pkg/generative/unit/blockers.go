package unit

import "testing"

func blockers(
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
