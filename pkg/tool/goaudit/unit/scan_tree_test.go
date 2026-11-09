package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/unit/scan_tester"
	"testing"
)

func TestTreeSkipsUnderscoreDirectories(t *testing.T) {
	base := t.TempDir()
	scan_tester.WriteSource(base, "pkg/alfa/alfa_test.go")
	scan_tester.WriteSource(base, "tmp/_bravo/bravo/bravo_test.go")
	scan_tester.WriteSource(base, "pkg/charlie/_probe/charlie_test.go")
	v := scan.Tree(base)
	assert.Integer(t, 1, len(v.Files()))
	result := scan.MisplacedTests(v)
	assert.Integer(t, 1, len(result))
	assert.String(t, "pkg/alfa/alfa_test.go", result[0].Path)
}
