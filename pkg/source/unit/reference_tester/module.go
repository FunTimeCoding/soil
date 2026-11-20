package reference_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
)

func (o *Tester) Module() string {
	o.t.Helper()
	directory := o.t.TempDir()
	system_tester.New(o.t).Write(
		directory,
		"alfa/server.go",
		"package alfa\n\ntype Server struct {\n\tPort int\n}\n\nfunc NewServer() *Server {\n\treturn &Server{}\n}\n\nfunc (s *Server) Start() {}\n\ntype Box[T any] struct {\n\tValue T\n}\n\nfunc (b *Box[T]) Get() T {\n\treturn b.Value\n}\n",
	)
	system_tester.New(o.t).Write(
		directory,
		"bravo/use.go",
		"package bravo\n\nimport \"example/alfa\"\n\nfunc Use() int {\n\ts := alfa.NewServer()\n\ts.Start()\n\tb := &alfa.Box[int]{}\n\n\treturn s.Port + b.Get()\n}\n",
	)
	system_tester.New(o.t).Write(
		directory,
		"bravo/use_test.go",
		"package bravo\n\nimport (\n\t\"example/alfa\"\n\t\"testing\"\n)\n\nfunc TestUse(t *testing.T) {\n\t_ = alfa.NewServer()\n}\n",
	)
	system_tester.New(o.t).Write(
		directory,
		"bravo/external_test.go",
		"package bravo_test\n\nimport (\n\t\"example/alfa\"\n\t\"testing\"\n)\n\nfunc TestExternal(t *testing.T) {\n\t_ = alfa.NewServer()\n}\n",
	)
	system_tester.New(o.t).Write(
		directory,
		"charlie/unit/only_test.go",
		"package unit\n\nimport (\n\t\"example/alfa\"\n\t\"testing\"\n)\n\nfunc TestOnly(t *testing.T) {\n\t_ = alfa.NewServer()\n}\n",
	)
	testutil.WriteModFile(o.t, directory, "example")

	return directory
}
