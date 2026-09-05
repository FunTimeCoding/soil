package route_guard

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/route_guard"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestBlocked(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	p.PkgPath = "example.test/pkg/tool/sample/web"
	route_guard.Check(p, results)
	testutil.AssertBlocked(t, results, 3)
}

func TestModelContextInScope(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	p.PkgPath = "example.test/pkg/tool/sample/model_context"
	route_guard.Check(p, results)
	testutil.AssertBlocked(t, results, 3)
}

func TestWebSubpackageInScope(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	p.PkgPath = "example.test/pkg/tool/sample/web/conversations"
	route_guard.Check(p, results)
	testutil.AssertBlocked(t, results, 3)
}

func TestOutOfScope(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	route_guard.Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}
