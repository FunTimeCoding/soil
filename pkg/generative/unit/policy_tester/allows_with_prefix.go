package policy_tester

import "testing"

func AllowsWithPrefix(
	t *testing.T,
	root string,
	scratchpad string,
	prefix string,
	command string,
) bool {
	t.Helper()

	return newWalk(t, root, scratchpad).
		WithScratchPrefixes(prefix).
		Allows(command)
}
