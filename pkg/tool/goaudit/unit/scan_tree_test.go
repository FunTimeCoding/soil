package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/warm_tree"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/unit/scan_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestTreeSkipsUnderscoreDirectories(t *testing.T) {
	base := t.TempDir()
	scan_tester.WriteSource(base, "pkg/alfa/alfa_test.go")
	scan_tester.WriteSource(base, "scratch/_bravo/bravo/bravo_test.go")
	scan_tester.WriteSource(base, "pkg/charlie/_probe/charlie_test.go")
	v := scan.Tree(warm_tree.New(base))
	assert.Integer(t, 1, len(v.Files()))
	result := scan.MisplacedTests(v)
	assert.Integer(t, 1, len(result))
	assert.String(t, "pkg/alfa/alfa_test.go", result[0].Path)
}

func TestTreeSkipsTemporaryDirectories(t *testing.T) {
	base := t.TempDir()
	scan_tester.WriteSource(base, "pkg/alfa/alfa_test.go")
	scan_tester.WriteSource(base, "tmp/bravo/bravo_test.go")
	scan_tester.WriteSource(base, "pkg/charlie/tmp/charlie_test.go")
	assert.Strings(
		t,
		[]string{"pkg/alfa/alfa_test.go"},
		scan.Tree(warm_tree.New(base)).Files(),
	)
}

func TestTreeStopsAtNestedModule(t *testing.T) {
	base := t.TempDir()
	scan_tester.WriteSource(base, "pkg/alfa/alfa_test.go")
	testutil.WriteFile(t, base, "graft/bravo/go.mod", "module bravo\n")
	scan_tester.WriteSource(base, "graft/bravo/bravo_test.go")
	assert.Strings(
		t,
		[]string{"pkg/alfa/alfa_test.go"},
		scan.Tree(warm_tree.New(base)).Files(),
	)
}

func TestTreeSkipsLinkedDirectories(t *testing.T) {
	base := t.TempDir()
	scan_tester.WriteSource(base, "pkg/alfa/alfa_test.go")
	assert.FatalOnError(
		t,
		os.Symlink(
			filepath.Join(base, constant.PackageDirectory),
			filepath.Join(base, "linked"),
		),
	)
	assert.Strings(
		t,
		[]string{"pkg/alfa/alfa_test.go"},
		scan.Tree(warm_tree.New(base)).Files(),
	)
}
