package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/integration/check_tester"
	"testing"
)

func TestSecondRunChecksOnlyPackageWithFindings(t *testing.T) {
	c := check_tester.New(t)
	root := c.CacheModule()
	index := t.TempDir()
	first, firstPrinted := c.CachedRun(root, index, false)
	assert.Strings(
		t,
		[]string{"pkg/alfa", "pkg/bravo", "pkg/response"},
		first.Checked,
	)
	second, secondPrinted := c.CachedRun(root, index, false)
	assert.Strings(t, []string{"pkg/response"}, second.Checked)
	assert.String(t, firstPrinted, secondPrinted)
}

func TestInventoryChangeChecksCachedPackagesAgain(t *testing.T) {
	c := check_tester.New(t)
	root := c.CacheModule()
	index := t.TempDir()
	other := t.TempDir()
	testutil.WriteModFile(t, other, "other.test")
	alone := inventory.New()
	alone.Add("example", root)
	c.InventoryRun(root, index, alone)
	assert.Strings(
		t,
		[]string{"pkg/response"},
		c.InventoryRun(root, index, alone).Checked,
	)
	wider := inventory.New()
	wider.Add("example", root)
	wider.Add("other", other)
	assert.Strings(
		t,
		[]string{"pkg/alfa", "pkg/bravo", "pkg/response"},
		c.InventoryRun(root, index, wider).Checked,
	)
}

func TestEditCheckedAgainAloneWhenSummaryHolds(t *testing.T) {
	c := check_tester.New(t)
	root := c.CacheModule()
	index := t.TempDir()
	c.CachedRun(root, index, false)
	system_tester.New(t).Write(
		root,
		"pkg/alfa/alfa.go",
		"package alfa\n\nfunc Alfa() int {\n\treturn 2\n}\n",
	)
	second, _ := c.CachedRun(root, index, false)
	assert.Strings(t, []string{"pkg/alfa", "pkg/response"}, second.Checked)
}

func TestSummaryChangeChecksImporterAgain(t *testing.T) {
	c := check_tester.New(t)
	root := c.CacheModule()
	system_tester.New(t).Write(
		root,
		"pkg/alfa/alfa.go",
		"package alfa\n\nimport \"io\"\n\nfunc Alfa(c io.Closer) int {\n\treturn 1\n}\n",
	)
	system_tester.New(t).Write(
		root,
		"pkg/bravo/bravo.go",
		"package bravo\n\nimport \"example/pkg/alfa\"\n\nfunc Bravo() int {\n\treturn alfa.Alfa(nil)\n}\n",
	)
	index := t.TempDir()
	c.CachedRun(root, index, false)
	system_tester.New(t).Write(
		root,
		"pkg/alfa/alfa.go",
		"package alfa\n\nimport \"io\"\n\nfunc Alfa(c io.Closer) int {\n\t_ = c.Close()\n\n\treturn 1\n}\n",
	)
	second, printed := c.CachedRun(root, index, false)
	assert.Strings(
		t,
		[]string{"pkg/alfa", "pkg/bravo", "pkg/response"},
		second.Checked,
	)
	assert.StringNotContains(t, "pkg/alfa/alfa.go", printed)
}

func TestFreshChecksEverything(t *testing.T) {
	c := check_tester.New(t)
	root := c.CacheModule()
	index := t.TempDir()
	c.CachedRun(root, index, false)
	second, _ := c.CachedRun(root, index, true)
	assert.Strings(
		t,
		[]string{"pkg/alfa", "pkg/bravo", "pkg/response"},
		second.Checked,
	)
}

func TestNewConstantRevealsLiteralInCachedPackage(t *testing.T) {
	c := check_tester.New(t)
	root := c.CacheModule()
	index := t.TempDir()
	c.CachedRun(root, index, false)
	system_tester.New(t).Write(
		root,
		"pkg/constant/constant.go",
		"package constant\n\nconst Delta = \"delta echo\"\n",
	)
	_, printed := c.CachedRun(root, index, false)
	assert.StringContains(
		t,
		`pkg/bravo/bravo.go:6: string literal "delta echo" has constant constant.Delta`,
		printed,
	)
}

func TestCommentSettingRevealsCommentInCachedPackage(t *testing.T) {
	c := check_tester.New(t)
	root := c.CacheModule()
	system_tester.New(t).Write(
		root,
		"pkg/alfa/alfa.go",
		"// Package alfa returns one.\npackage alfa\n\nfunc Alfa() int {\n\treturn 1\n}\n",
	)
	index := t.TempDir()
	_, before := c.CachedRun(root, index, false)
	assert.StringNotContains(t, "pkg/alfa/alfa.go", before)
	system_tester.New(t).Write(
		root,
		"strata/tool/goanalyze.yaml",
		"comment: true\n",
	)
	_, after := c.CachedRun(root, index, false)
	assert.StringContains(t, "pkg/alfa/alfa.go:1", after)
}

func TestVanishedInterfaceRevealsNamingInCachedPackage(t *testing.T) {
	c := check_tester.New(t)
	root := c.CacheModule()
	system_tester.New(t).Write(
		root,
		"pkg/foxtrot/foxtrot.go",
		"package foxtrot\n\ntype Foxtrot struct{}\n",
	)
	system_tester.New(t).Write(
		root,
		"pkg/foxtrot/url.go",
		"package foxtrot\n\nfunc (f *Foxtrot) URL() string {\n\treturn \"golf\"\n}\n",
	)
	system_tester.New(t).Write(
		root,
		"pkg/hotel/hotel.go",
		"package hotel\n\ntype Hotel interface {\n\tURL() string\n}\n",
	)
	index := t.TempDir()
	_, before := c.CachedRun(root, index, false)
	assert.StringNotContains(t, "pkg/foxtrot", before)
	system_tester.New(t).Write(root, "pkg/hotel/hotel.go", "package hotel\n")
	_, after := c.CachedRun(root, index, false)
	assert.StringContains(t, "pkg/foxtrot/url.go: renamed URL", after)
}
