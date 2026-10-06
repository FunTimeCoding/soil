package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"os"
	"path/filepath"
	"testing"
)

func TestUpdateFollowsABodyEdit(t *testing.T) {
	_, user := replacedModules(t)
	updateMatchesBuild(
		t,
		user,
		func() {
			testutil.WriteFile(
				t,
				user,
				"user/user.go",
				"package user\n\nimport \"other.test/lib\"\n\nfunc Use() string {\n\treturn lib.Name() + \"!\"\n}\n",
			)
		},
	)
}

func TestUpdateFollowsANewPackageAndImport(t *testing.T) {
	_, user := replacedModules(t)
	g := updateMatchesBuild(
		t,
		user,
		func() {
			testutil.WriteFile(
				t,
				user,
				"extra/extra.go",
				"package extra\n\nfunc Extra() string {\n\treturn \"extra\"\n}\n",
			)
			testutil.WriteFile(
				t,
				user,
				"user/user.go",
				"package user\n\nimport (\n\t\"example/extra\"\n\t\"other.test/lib\"\n)\n\nfunc Use() string {\n\treturn lib.Name() + extra.Extra()\n}\n",
			)
		},
	)
	assert.NotNil(t, g.Nodes["example/extra"])
}

func TestUpdateFollowsADeletedPackage(t *testing.T) {
	_, user := replacedModules(t)
	g := updateMatchesBuild(
		t,
		user,
		func() {
			assert.FatalOnError(t, os.RemoveAll(filepath.Join(user, "user")))
		},
	)
	assert.Nil(t, g.Nodes["example/user"])
	assert.Nil(t, g.Nodes["other.test/lib"])
}

func TestUpdateFollowsATestOnlyDirectory(t *testing.T) {
	_, user := replacedModules(t)
	updateMatchesBuild(
		t,
		user,
		func() {
			testutil.WriteFile(
				t,
				user,
				"charlie/unit/only_test.go",
				"package unit\n\nimport (\n\t\"example/user\"\n\t\"testing\"\n)\n\nfunc TestOnly(t *testing.T) {\n\t_ = user.Use()\n}\n",
			)
		},
	)
}

func TestUpdateRebuildsWhenTheTagsChange(t *testing.T) {
	_, user := replacedModules(t)
	tagged := "//go:build custom\n\npackage user\n\nfunc Tagged() string {\n\treturn \"tagged\"\n}\n"
	g := updateMatchesBuild(
		t,
		user,
		func() {
			testutil.WriteFile(t, user, "user/tagged.go", tagged)
		},
	)
	assert.Strings(t, []string{"custom"}, g.Tags)
	updateMatchesBuild(
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

func TestUpdateFollowsAReplacedModule(t *testing.T) {
	library, user := replacedModules(t)
	testutil.WriteFile(
		t,
		library,
		"q/q.go",
		"package q\n\nfunc Q() string {\n\treturn \"q\"\n}\n",
	)
	g := updateMatchesBuild(
		t,
		user,
		func() {
			testutil.WriteFile(
				t,
				library,
				"lib.go",
				"package lib\n\nimport \"other.test/lib/q\"\n\nfunc Name() string {\n\treturn q.Q()\n}\n",
			)
		},
	)
	assert.NotNil(t, g.Nodes["other.test/lib/q"])
	g = updateMatchesBuild(
		t,
		user,
		func() {
			testutil.WriteFile(
				t,
				library,
				"lib.go",
				"package lib\n\nfunc Name() string {\n\treturn \"lib\"\n}\n",
			)
		},
	)
	assert.Nil(t, g.Nodes["other.test/lib/q"])
}

func TestUpdateKeepsWhatAnUnchangedReplacedPackageReaches(t *testing.T) {
	library, user := replacedModules(t)
	testutil.WriteFile(
		t,
		library,
		"q/q.go",
		"package q\n\nfunc Q() string {\n\treturn \"q\"\n}\n",
	)
	testutil.WriteFile(
		t,
		library,
		"lib.go",
		"package lib\n\nimport \"other.test/lib/q\"\n\nfunc Name() string {\n\treturn q.Q()\n}\n",
	)
	g := updateMatchesBuild(
		t,
		user,
		func() {
			testutil.WriteFile(
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
	library, user := replacedModules(t)
	updateMatchesBuild(
		t,
		user,
		func() {
			testutil.WriteFile(
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

func TestUpdateIgnoresWhatTheModuleDoesNotBuild(t *testing.T) {
	_, user := replacedModules(t)
	updateMatchesBuild(
		t,
		user,
		func() {
			testutil.WriteFile(
				t,
				user,
				"user/testdata/fixture.go",
				"package fixture\n",
			)
		},
	)
}
