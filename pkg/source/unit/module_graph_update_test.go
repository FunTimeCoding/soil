package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/unit/module_graph_tester"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestUpdateFollowsBodyEdit(t *testing.T) {
	_, user := module_graph_tester.Replaced(t)
	module_graph_tester.Updated(
		t,
		user,
		func() {
			system_tester.Write(
				t,
				user,
				"user/user.go",
				"package user\n\nimport \"other.test/lib\"\n\nfunc Use() string {\n\treturn lib.Name() + \"!\"\n}\n",
			)
		},
	)
}

func TestUpdateFollowsNewPackageAndImport(t *testing.T) {
	_, user := module_graph_tester.Replaced(t)
	g := module_graph_tester.Updated(
		t,
		user,
		func() {
			system_tester.Write(
				t,
				user,
				"extra/extra.go",
				"package extra\n\nfunc Extra() string {\n\treturn \"extra\"\n}\n",
			)
			system_tester.Write(
				t,
				user,
				"user/user.go",
				"package user\n\nimport (\n\t\"example/extra\"\n\t\"other.test/lib\"\n)\n\nfunc Use() string {\n\treturn lib.Name() + extra.Extra()\n}\n",
			)
		},
	)
	assert.NotNil(t, g.Nodes["example/extra"])
}

func TestUpdateFollowsDeletedPackage(t *testing.T) {
	_, user := module_graph_tester.Replaced(t)
	g := module_graph_tester.Updated(
		t,
		user,
		func() {
			assert.FatalOnError(t, os.RemoveAll(filepath.Join(user, "user")))
		},
	)
	assert.Nil(t, g.Nodes["example/user"])
	assert.Nil(t, g.Nodes["other.test/lib"])
}

func TestUpdateFollowsTestOnlyDirectory(t *testing.T) {
	_, user := module_graph_tester.Replaced(t)
	module_graph_tester.Updated(
		t,
		user,
		func() {
			system_tester.Write(
				t,
				user,
				"charlie/unit/only_test.go",
				"package unit\n\nimport (\n\t\"example/user\"\n\t\"testing\"\n)\n\nfunc TestOnly(t *testing.T) {\n\t_ = user.Use()\n}\n",
			)
		},
	)
}

func TestUpdateRebuildsWhenTagsChange(t *testing.T) {
	_, user := module_graph_tester.Replaced(t)
	tagged := "//go:build custom\n\npackage user\n\nfunc Tagged() string {\n\treturn \"tagged\"\n}\n"
	g := module_graph_tester.Updated(
		t,
		user,
		func() {
			system_tester.Write(t, user, "user/tagged.go", tagged)
		},
	)
	assert.Strings(t, []string{"custom"}, g.Tags)
	module_graph_tester.Updated(
		t,
		user,
		func() {
			assert.FatalOnError(
				t,
				os.Remove(filepath.Join(user, "user/tagged.go")),
			)
		},
	)
}

func TestUpdateFollowsReplacedModule(t *testing.T) {
	library, user := module_graph_tester.Replaced(t)
	system_tester.Write(
		t,
		library,
		"q/q.go",
		"package q\n\nfunc Q() string {\n\treturn \"q\"\n}\n",
	)
	g := module_graph_tester.Updated(
		t,
		user,
		func() {
			system_tester.Write(
				t,
				library,
				"lib.go",
				"package lib\n\nimport \"other.test/lib/q\"\n\nfunc Name() string {\n\treturn q.Q()\n}\n",
			)
		},
	)
	assert.NotNil(t, g.Nodes["other.test/lib/q"])
	g = module_graph_tester.Updated(
		t,
		user,
		func() {
			system_tester.Write(
				t,
				library,
				"lib.go",
				"package lib\n\nfunc Name() string {\n\treturn \"lib\"\n}\n",
			)
		},
	)
	assert.Nil(t, g.Nodes["other.test/lib/q"])
}

func TestUpdateKeepsWhatUnchangedReplacedPackageReaches(t *testing.T) {
	library, user := module_graph_tester.Replaced(t)
	system_tester.Write(
		t,
		library,
		"q/q.go",
		"package q\n\nfunc Q() string {\n\treturn \"q\"\n}\n",
	)
	system_tester.Write(
		t,
		library,
		"lib.go",
		"package lib\n\nimport \"other.test/lib/q\"\n\nfunc Name() string {\n\treturn q.Q()\n}\n",
	)
	g := module_graph_tester.Updated(
		t,
		user,
		func() {
			system_tester.Write(
				t,
				user,
				"user/more.go",
				"package user\n\nfunc More() string {\n\treturn \"more\"\n}\n",
			)
		},
	)
	assert.NotNil(t, g.Nodes["other.test/lib/q"])
}

func TestUpdateRebuildsWhenGoModChanges(t *testing.T) {
	library, user := module_graph_tester.Replaced(t)
	module_graph_tester.Updated(
		t,
		user,
		func() {
			system_tester.Write(
				t,
				user,
				"go.mod",
				fmt.Sprintf(
					"module example\n\ngo 1.22\n\nrequire other.test/lib v0.0.0\n\nreplace other.test/lib => %s\n\n// edited\n",
					library,
				),
			)
		},
	)
}

func TestUpdateIgnoresWhatModuleDoesNotBuild(t *testing.T) {
	_, user := module_graph_tester.Replaced(t)
	module_graph_tester.Updated(
		t,
		user,
		func() {
			system_tester.Write(
				t,
				user,
				"user/testdata/fixture.go",
				"package fixture\n",
			)
		},
	)
}
