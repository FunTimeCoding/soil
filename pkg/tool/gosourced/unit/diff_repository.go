package unit

import (
	"github.com/funtimecoding/soil/pkg/git/unit/repository_tester"
	"testing"
)

func diffRepository(t *testing.T) *repository_tester.Tester {
	t.Helper()
	r := repository_tester.New(t)
	r.Write("go.mod", "module example.org/root\n")
	r.Write(
		"alfa.go",
		"package alfa\n\nfunc Bravo() int {\n\treturn 1\n}\n\nfunc Charlie() int {\n\treturn 3\n}\n",
	)
	r.Commit("alfa")
	r.Write(
		"alfa.go",
		"package alfa\n\nfunc Bravo() int {\n\treturn 2\n}\n\nfunc Charlie() int {\n\treturn 3\n}\n",
	)
	repository_tester.Command(r.Clone, "add", "alfa.go")

	return r
}
