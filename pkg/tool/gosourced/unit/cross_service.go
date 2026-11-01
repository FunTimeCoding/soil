package unit

import (
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/git/unit/repository_tester"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/source/inventory/module"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"testing"
)

func crossService(t *testing.T) *service.Service {
	t.Helper()
	l := repository_tester.New(t)
	l.Write("go.mod", "module example.org/lib\n")
	l.Write(
		"charlie/charlie.go",
		"package charlie\n\nfunc Delta() int {\n\treturn 4\n}\n",
	)
	l.Commit("delta")
	repository_tester.Command(l.Clone, gitConstant.Tag, "v0.1.0")
	l.Write(
		"charlie/bravo.go",
		"package charlie\n\nfunc Bravo() int {\n\treturn 2\n}\n",
	)
	l.Write(
		"charlie/charlie.go",
		"package charlie\n\nfunc Delta() int {\n\treturn 5\n}\n",
	)
	l.Commit("bravo")
	repository_tester.Command(l.Clone, gitConstant.Tag, "v0.2.0")
	u := repository_tester.New(t)
	u.Write(
		"go.mod",
		"module example.org/user\n\nrequire example.org/lib v0.1.0\n",
	)
	u.Write(
		"alfa/alfa.go",
		"package alfa\n\nfunc Bravo() int {\n\treturn 2\n}\n",
	)
	u.Write(
		"echo/echo.go",
		"package echo\n\nimport \"example.org/user/alfa\"\n\nfunc Echo() int {\n\treturn alfa.Bravo()\n}\n",
	)
	u.Commit("alfa")
	repository_tester.Command(u.Clone, "rm", "-q", "alfa/alfa.go")
	u.Write(
		"go.mod",
		"module example.org/user\n\nrequire example.org/lib v0.2.0\n",
	)
	u.Write(
		"echo/echo.go",
		"package echo\n\nimport \"example.org/lib/charlie\"\n\nfunc Echo() int {\n\treturn charlie.Bravo()\n}\n",
	)
	u.Commit("promote")
	i := inventory.New()
	i.Modules = []module.Module{
		{Name: "user", Directory: u.Clone},
		{Name: "lib", Directory: l.Clone},
	}

	return service.New(i)
}
