package repository_tester

import "testing"

func Modified(t *testing.T) *Tester {
	t.Helper()
	r := New(t)
	r.Write("go.mod", "module alfa\n")
	r.Write("graft/bravo/go.mod", "module bravo\n")
	r.Write("graft/bravo/seed.md", "# Seed\n")
	r.Commit("modules")
	r.Write("go.mod", "module alfa\n\ngo 1.27\n")
	r.Write("graft/bravo/go.mod", "module bravo\n\ngo 1.27\n")
	r.Write("graft/bravo/staged.md", "# Staged\n")
	Command(r.Clone, "add", "graft/bravo/staged.md")
	r.Write("graft/bravo/untracked.md", "# Untracked\n")

	return r
}
