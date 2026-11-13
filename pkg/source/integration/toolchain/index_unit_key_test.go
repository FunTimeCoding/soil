package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/unit/kind_tester"
	"testing"
)

func TestUnitKeysHoldWhenUnrelatedPackageChanges(t *testing.T) {
	directory := keyedModule(t)
	store := t.TempDir()
	declared := kind_tester.Declared("golf")
	unit, closure := alfaKeys(
		t,
		index.New(store, directory, declared),
		declared,
	)
	testutil.WriteFile(
		t,
		directory,
		"delta/delta.go",
		"package delta\n\nfunc Delta() int {\n\treturn 5\n}\n\nfunc Hotel() {}\n",
	)
	after, afterClosure := alfaKeys(
		t,
		index.New(store, directory, declared),
		declared,
	)
	assert.String(t, unit, after)
	assert.String(t, closure, afterClosure)
}

func TestUnitKeyMovesWithTestFile(t *testing.T) {
	directory := keyedModule(t)
	store := t.TempDir()
	declared := kind_tester.Declared("golf")
	unit, _ := alfaKeys(t, index.New(store, directory, declared), declared)
	testutil.WriteFile(
		t,
		directory,
		"alfa/alfa_test.go",
		"package alfa\n\nimport \"testing\"\n\nfunc TestAlfa(t *testing.T) {\n\tAlfa()\n}\n",
	)
	after, _ := alfaKeys(t, index.New(store, directory, declared), declared)
	assert.True(t, unit != after)
}

func TestUnitKeyMovesWithImportedAPI(t *testing.T) {
	directory := keyedModule(t)
	store := t.TempDir()
	declared := kind_tester.Declared("golf")
	unit, _ := alfaKeys(t, index.New(store, directory, declared), declared)
	testutil.WriteFile(
		t,
		directory,
		"bravo/bravo.go",
		"package bravo\n\nimport \"example/charlie\"\n\nfunc Bravo() int {\n\treturn charlie.Charlie()\n}\n\nfunc Echo() {}\n",
	)
	after, _ := alfaKeys(t, index.New(store, directory, declared), declared)
	assert.True(t, unit != after)
}

func TestKeysHoldWhenBodyDeepInClosureChanges(t *testing.T) {
	directory := keyedModule(t)
	store := t.TempDir()
	declared := kind_tester.Declared("golf")
	unit, closure := alfaKeys(
		t,
		index.New(store, directory, declared),
		declared,
	)
	testutil.WriteFile(
		t,
		directory,
		"charlie/charlie.go",
		"package charlie\n\nfunc Charlie() int {\n\treturn 2\n}\n",
	)
	after, afterClosure := alfaKeys(
		t,
		index.New(store, directory, declared),
		declared,
	)
	assert.String(t, unit, after)
	assert.String(t, closure, afterClosure)
}

func TestOnlyClosureKeyMovesWithFactDeepInClosure(t *testing.T) {
	directory := keyedModule(t)
	store := t.TempDir()
	declared := kind_tester.Declared("golf")
	unit, closure := alfaKeys(
		t,
		index.New(store, directory, declared),
		declared,
	)
	testutil.WriteFile(
		t,
		directory,
		"charlie/charlie.go",
		"package charlie\n\nfunc Charlie() int {\n\treturn 1\n}\n\nfunc Echo() {}\n",
	)
	after, afterClosure := alfaKeys(
		t,
		index.New(store, directory, declared),
		declared,
	)
	assert.String(t, unit, after)
	assert.True(t, closure != afterClosure)
}

func TestClosureKeySameWhetherFactsFreshOrStored(t *testing.T) {
	directory := keyedModule(t)
	store := t.TempDir()
	listed := kind_tester.Listed("golf")
	_, fresh := alfaKeys(t, index.New(store, directory, listed), listed)
	_, stored := alfaKeys(t, index.New(store, directory, listed), listed)
	assert.String(t, fresh, stored)
}

func TestUnitKeyMovesWithExtractingBuild(t *testing.T) {
	directory := keyedModule(t)
	store := t.TempDir()
	declared := kind_tester.Declared("golf")
	first, _ := alfaKeys(
		t,
		index.New(store, directory, declared).WithFingerprint("first"),
		declared,
	)
	second, _ := alfaKeys(
		t,
		index.New(store, directory, declared).WithFingerprint("second"),
		declared,
	)
	assert.True(t, first != second)
}

func TestNoUnitKeysWhenImportDoesNotLoad(t *testing.T) {
	directory := keyedModule(t)
	testutil.WriteFile(
		t,
		directory,
		"bravo/bravo.go",
		"package bravo\n\nimport \"example/charlie\"\n\nfunc Bravo() int {\n\treturn charlie.Charlie(\n}\n",
	)
	declared := kind_tester.Declared("golf")
	w := index.New(t.TempDir(), directory, declared)
	_, okay := w.UnitKey("example/alfa")
	assert.False(t, okay)
	_, okay = w.ClosureKey("example/alfa", declared)
	assert.False(t, okay)
}
