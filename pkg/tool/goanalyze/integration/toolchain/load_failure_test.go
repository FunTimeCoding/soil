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

func TestBrokenAnalyzedPackageAnswersThatItDoesNotLoad(t *testing.T) {
	if o := scope_tester.ChildOptions("./caller/..."); o != nil {
		goanalyze.Run(o)

		return
	}

	s := scope_tester.New(t)

	for _, full := range []bool{false, true} {
		directory := t.TempDir()
		s.HelperModule(directory, "example")
		s.Caller(directory, "example")
		s.BrokenPackage(directory, "caller")
		r := s.Child(directory, full)
		assert.Integer(t, 1, r.Exit)
		assert.StringContains(
			t,
			"does not load: example/caller - ",
			r.ErrorString,
		)
		assert.StringNotContains(t, "panic", r.ErrorString)
	}
}

func TestBrokenReachedPackageIsSkippedWithNotice(t *testing.T) {
	s := scope_tester.New(t)
	root := t.TempDir()
	helper := filepath.Join(root, "helper_module")
	application := filepath.Join(root, "application")
	s.HelperModule(helper, "helper.example")
	s.BrokenPackage(helper, "helper")
	s.Caller(application, "helper.example")
	testutil.WriteModFile(
		t,
		application,
		"application.example",
		"require helper.example v0.0.0",
		"replace helper.example => ../helper_module",
	)
	var notice bytes.Buffer
	o := s.Options(application, true, "./caller/...")
	o.Notice = &notice
	testutil.AssertBlockedContains(t, goanalyze.Analyze(o), "s is never closed")
	assert.String(
		t,
		"facts skipped: helper.example/helper does not load\n",
		notice.String(),
	)
}
