package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocheck"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/integration/check_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	gosourcedConstant "github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/base"
	"path/filepath"
	"strconv"
	"testing"
)

func TestDelegatedCheckMatchesLocalRun(t *testing.T) {
	c := check_tester.New(t)
	root := c.StackModule()
	c.ServeSource(base.WithModule("example", root))
	local := option.New()
	local.Root = root
	local.Index = t.TempDir()
	expected, e := gocheck.Run(local)
	assert.FatalOnError(t, e)
	remote := option.New()
	remote.Root = root
	actual, delegated, f := gocheck.Delegate(remote)
	assert.FatalOnError(t, f)
	assert.True(t, delegated)
	assert.True(t, len(actual.Sections) > 0)
	assert.True(t, len(actual.Changes) > 0)
	assert.Exported(t, expected, actual)
}

func TestDelegatedFixCommitsThroughDaemon(t *testing.T) {
	c := check_tester.New(t)
	root := c.StackModule()
	c.ServeSource(base.WithModule("example", root))
	o := option.New()
	o.Root = root
	o.Fix = true
	actual, delegated, e := gocheck.Delegate(o)
	assert.FatalOnError(t, e)
	assert.True(t, delegated)
	assert.Count(t, 0, actual.Changes)
	assert.String(
		t,
		"package spacing\n\nfunc Alfa() int {\n\tbravo := 1\n\n\treturn bravo\n}\n",
		testutil.ReadFile(t, filepath.Join(root, "pkg/spacing/alfa.go")),
	)
}

func TestDelegatedCheckSeesChangeBetweenChecks(t *testing.T) {
	c := check_tester.New(t)
	root := c.StackModule()
	c.ServeSource(base.WithModule("example", root))
	first := option.New()
	first.Root = root
	before, _, e := gocheck.Delegate(first)
	assert.FatalOnError(t, e)
	assert.True(t, check_tester.Changes(before, "pkg/spacing/alfa.go"))
	system_tester.New(t).Write(
		root,
		"pkg/spacing/alfa.go",
		"package spacing\n\nfunc Alfa() int {\n\tbravo := 1\n\n\treturn bravo\n}\n",
	)
	second := option.New()
	second.Root = root
	after, _, f := gocheck.Delegate(second)
	assert.FatalOnError(t, f)
	assert.False(t, check_tester.Changes(after, "pkg/spacing/alfa.go"))
}

func TestDelegateDeclinesRootOutsideInventory(t *testing.T) {
	c := check_tester.New(t)
	c.ServeSource()
	o := option.New()
	o.Root = c.StackModule()
	_, delegated, e := gocheck.Delegate(o)
	assert.FatalOnError(t, e)
	assert.False(t, delegated)
}

func TestDelegateDeclinesWithoutDaemon(t *testing.T) {
	t.Setenv(gosourcedConstant.HostEnvironment, "localhost")
	t.Setenv(
		gosourcedConstant.PortEnvironment,
		strconv.Itoa(system.FindUnusedPort(20000)),
	)
	t.Setenv(gosourcedConstant.InsecureEnvironment, "true")
	t.Setenv(gosourcedConstant.TokenEnvironment, "alfa")
	o := option.New()
	o.Root = t.TempDir()
	_, delegated, e := gocheck.Delegate(o)
	assert.FatalOnError(t, e)
	assert.False(t, delegated)
}
