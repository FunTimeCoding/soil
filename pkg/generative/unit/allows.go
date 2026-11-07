package unit

import "testing"

func allows(
	t *testing.T,
	root string,
	scratchpad string,
	command string,
) bool {
	t.Helper()

	return newWalk(t, root, scratchpad).Allows(command)
}
