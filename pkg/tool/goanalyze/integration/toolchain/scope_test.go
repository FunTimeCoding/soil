package toolchain

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/unit/scope_tester"
	"path/filepath"
	"testing"
)

func TestScopedRunSeesHelperOutsidePatterns(t *testing.T) {
	s := scope_tester.New(t)

	for _, full := range []bool{false, true} {
		directory := t.TempDir()
		s.HelperModule(directory, "example")
		s.Caller(directory, "example")
		testutil.AssertNotBlockedContains(
			t,
			goanalyze.Analyze(s.Options(directory, full, "./caller/...")),
			"never closed",
		)
	}
}

func TestFullRunReportsHelperLeak(t *testing.T) {
	s := scope_tester.New(t)

	for _, full := range []bool{false, true} {
		directory := t.TempDir()
		s.HelperModule(directory, "example")
		s.Caller(directory, "example")
		testutil.AssertBlockedContains(
			t,
			goanalyze.Analyze(s.Options(directory, full)),
			"r is never closed",
		)
	}
}

func TestReplacedModuleHelperIsIndexed(t *testing.T) {
	s := scope_tester.New(t)

	for _, full := range []bool{false, true} {
		root := t.TempDir()
		helper := filepath.Join(root, "helper_module")
		application := filepath.Join(root, "application")
		s.HelperModule(helper, "helper.example")
		s.Caller(application, "helper.example")
		testutil.WriteModFile(
			t,
			application,
			"application.example",
			"require helper.example v0.0.0",
			"replace helper.example => ../helper_module",
		)
		testutil.AssertNotBlockedContains(
			t,
			goanalyze.Analyze(s.Options(application, full, "./caller/...")),
			"never closed",
		)
	}
}

func TestCachedFactsFollowHelperBody(t *testing.T) {
	s := scope_tester.New(t)
	directory := t.TempDir()
	s.HelperModule(directory, "example")
	s.Caller(directory, "example")
	o := s.Options(directory, false, "./caller/...")
	testutil.AssertNotBlockedContains(t, goanalyze.Analyze(o), "never closed")
	testutil.AssertNotBlockedContains(t, goanalyze.Analyze(o), "never closed")
	s.HelperWithoutCleanup(directory)
	testutil.AssertBlockedContains(t, goanalyze.Analyze(o), "s is never closed")
	s.HelperModule(directory, "example")
	testutil.AssertNotBlockedContains(t, goanalyze.Analyze(o), "never closed")
}
