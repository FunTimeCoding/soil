package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/document_cache"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"strings"
	"testing"
)

func TestDocumentCacheRunsOncePerContent(t *testing.T) {
	c := document_cache.New()
	runs := 0
	run := func() *file_report.Report {
		runs++
		r := file_report.New("alfa.go", strings.NewReader("bravo\n"))
		r.AddConcern("charlie", "delta", "alfa.go", 1, "bravo", false)

		return r.Finalize()
	}
	c.Result("echo", "alfa.go", "bravo\n", run)
	c.Result("echo", "alfa.go", "bravo\n", run)
	assert.Integer(t, 1, runs)
	c.Result("echo", "alfa.go", "brave\n", run)
	c.Result("golf", "alfa.go", "bravo\n", run)
	assert.Integer(t, 3, runs)
}

func TestDocumentCacheHandsOutIndependentConcerns(t *testing.T) {
	c := document_cache.New()
	run := func() *file_report.Report {
		r := file_report.New("alfa.go", strings.NewReader("bravo\n"))
		r.AddConcern("charlie", "delta", "/root/alfa.go", 1, "bravo", true)

		return r.Finalize()
	}
	c.Result("echo", "alfa.go", "bravo\n", run)
	first := c.Result("echo", "alfa.go", "bravo\n", run)
	first.Concerns[0].Path = "alfa.go"
	first.Concerns[0].Fixed = false
	second := c.Result("echo", "alfa.go", "bravo\n", run)
	assert.String(t, "/root/alfa.go", second.Concerns[0].Path)
	assert.True(t, second.Concerns[0].Fixed)
}

func TestDocumentCacheDropsWhatRotationLeftUnused(t *testing.T) {
	c := document_cache.New()
	runs := 0
	run := func() *file_report.Report {
		runs++

		return file_report.New("alfa.go", strings.NewReader("")).Finalize()
	}
	c.Result("echo", "alfa.go", "bravo\n", run)
	c.Rotate()
	c.Result("echo", "alfa.go", "bravo\n", run)
	assert.Integer(t, 1, runs)
	c.Rotate()
	c.Rotate()
	c.Result("echo", "alfa.go", "bravo\n", run)
	assert.Integer(t, 2, runs)
}

func TestReportWithSharedDocumentsSeesEditedFile(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(
		t,
		root,
		"alfa/alfa.go",
		"package alfa\n\nfunc Alfa() int {\n\tbravo := 1\n\treturn bravo\n}\n",
	)
	first := option.New("", false)
	before := lint.Report("golint", root, first)
	assert.True(t, len(before.Entries) > 0)
	again := option.New("", false)
	again.Documents = first.Documents
	assert.Integer(
		t,
		len(before.Entries),
		len(lint.Report("golint", root, again).Entries),
	)
	testutil.WriteFile(
		t,
		root,
		"alfa/alfa.go",
		"package alfa\n\nfunc Alfa() int {\n\tbravo := 1\n\n\treturn bravo\n}\n",
	)
	after := option.New("", false)
	after.Documents = first.Documents
	assert.Count(t, 0, lint.Report("golint", root, after).Entries)
}

func TestReportWithSharedDocumentsResolvesPointersAgain(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(
		t,
		root,
		"doc/alfa.md",
		"# Alfa\n\nSee `doc/bravo.md`.\n",
	)
	first := option.New("", false)
	assert.Count(t, 1, lint.Report("golint", root, first).Entries)
	testutil.WriteFile(t, root, "doc/bravo.md", "# Bravo\n")
	after := option.New("", false)
	after.Documents = first.Documents
	assert.Count(t, 0, lint.Report("golint", root, after).Entries)
}
