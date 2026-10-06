package unit

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/unit/scope_tester"
	"path/filepath"
	"testing"
)

func TestScopedRunSeesHelperOutsidePatterns(t *testing.T) {
	for _, full := range []bool{false, true} {
		directory := t.TempDir()
		scope_tester.HelperModule(t, directory, "example")
		scope_tester.Caller(t, directory, "example")
		testutil.AssertNotBlockedContains(
			t,
			goanalyze.Analyze(
				scope_tester.Options(t, directory, full, "./caller/..."),
			),
			"never closed",
		)
	}
}

func TestFullRunReportsHelperLeak(t *testing.T) {
	for _, full := range []bool{false, true} {
		directory := t.TempDir()
		scope_tester.HelperModule(t, directory, "example")
		scope_tester.Caller(t, directory, "example")
		testutil.AssertBlockedContains(
			t,
			goanalyze.Analyze(scope_tester.Options(t, directory, full)),
			"r is never closed",
		)
	}
}

func TestReplacedModuleHelperIsIndexed(t *testing.T) {
	for _, full := range []bool{false, true} {
		root := t.TempDir()
		helper := filepath.Join(root, "helper_module")
		application := filepath.Join(root, "application")
		scope_tester.HelperModule(t, helper, "helper.example")
		scope_tester.Caller(t, application, "helper.example")
		testutil.WriteModFile(
			t,
			application,
			"application.example",
			"require helper.example v0.0.0",
			"replace helper.example => ../helper_module",
		)
		testutil.AssertNotBlockedContains(
			t,
			goanalyze.Analyze(
				scope_tester.Options(t, application, full, "./caller/..."),
			),
			"never closed",
		)
	}
}

func TestCachedFactsFollowHelperBody(t *testing.T) {
	directory := t.TempDir()
	scope_tester.HelperModule(t, directory, "example")
	scope_tester.Caller(t, directory, "example")
	o := scope_tester.Options(t, directory, false, "./caller/...")
	testutil.AssertNotBlockedContains(t, goanalyze.Analyze(o), "never closed")
	testutil.AssertNotBlockedContains(t, goanalyze.Analyze(o), "never closed")
	scope_tester.HelperWithoutCleanup(t, directory)
	testutil.AssertBlockedContains(t, goanalyze.Analyze(o), "s is never closed")
	scope_tester.HelperModule(t, directory, "example")
	testutil.AssertNotBlockedContains(t, goanalyze.Analyze(o), "never closed")
}
