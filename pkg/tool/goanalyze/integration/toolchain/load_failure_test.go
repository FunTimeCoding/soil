package toolchain

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/unit/scope_tester"
	"path/filepath"
	"testing"
)

func TestABrokenAnalyzedPackageAnswersThatItDoesNotLoad(t *testing.T) {
	if o := scope_tester.ChildOptions("./caller/..."); o != nil {
		goanalyze.Run(o)

		return
	}

	for _, full := range []bool{false, true} {
		directory := t.TempDir()
		scope_tester.HelperModule(t, directory, "example")
		scope_tester.Caller(t, directory, "example")
		scope_tester.BrokenPackage(t, directory, "caller")
		r := scope_tester.Child(t, directory, full)
		assert.Integer(t, 1, r.Exit)
		assert.StringContains(
			t,
			"does not load: example/caller - ",
			r.ErrorString,
		)
		assert.StringNotContains(t, "panic", r.ErrorString)
	}
}

func TestABrokenReachedPackageIsSkippedWithANotice(t *testing.T) {
	root := t.TempDir()
	helper := filepath.Join(root, "helper_module")
	application := filepath.Join(root, "application")
	scope_tester.HelperModule(t, helper, "helper.example")
	scope_tester.BrokenPackage(t, helper, "helper")
	scope_tester.Caller(t, application, "helper.example")
	testutil.WriteModFile(
		t,
		application,
		"application.example",
		"require helper.example v0.0.0",
		"replace helper.example => ../helper_module",
	)
	var notice bytes.Buffer
	o := scope_tester.Options(t, application, true, "./caller/...")
	o.Notice = &notice
	testutil.AssertBlockedContains(t, goanalyze.Analyze(o), "s is never closed")
	assert.String(
		t,
		"facts skipped: helper.example/helper does not load\n",
		notice.String(),
	)
}
