package policy_tester

import "testing"

func Allows(
	t *testing.T,
	root string,
	scratchpad string,
	command string,
) bool {
	t.Helper()

	return newWalk(t, root, scratchpad).Allows(command)
}
