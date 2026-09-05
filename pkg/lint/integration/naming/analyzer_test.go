package naming

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/naming"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestBlocked(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	naming.Check(p, results, face.New([]*packages.Package{p}))
	testutil.AssertBlocked(t, results, 57)
}
