package unit

import (
	"github.com/funtimecoding/soil/pkg/git/unit/repository_tester"
	"testing"
)

func modifiedRepository(t *testing.T) *repository_tester.Tester {
	t.Helper()
	r := repository_tester.New(t)
	r.Write("go.mod", "module alfa\n")
	r.Write("graft/bravo/go.mod", "module bravo\n")
	r.Write("graft/bravo/seed.md", "# Seed\n")
	r.Commit("modules")
	r.Write("go.mod", "module alfa\n\ngo 1.27\n")
	r.Write("graft/bravo/go.mod", "module bravo\n\ngo 1.27\n")
	r.Write("graft/bravo/staged.md", "# Staged\n")
	repository_tester.Command(r.Clone, "add", "graft/bravo/staged.md")
	r.Write("graft/bravo/untracked.md", "# Untracked\n")

	return r
}
