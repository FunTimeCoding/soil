package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gocheck"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"testing"
)

func TestFixModeAnalysisSkipsWhatFixFixed(t *testing.T) {
	root := stackModule(t)
	o := option.New()
	o.Root = root
	o.Fix = true
	o.Index = t.TempDir()
	printed := testutil.CaptureStdout(func() { gocheck.Print(checked(t, o)) })
	assert.StringNotContains(t, "pkg/format/delta.go:6: call exceeds", printed)
}

func TestReportModeMarksFindingOnRewrittenLine(t *testing.T) {
	root := stackModule(t)
	testutil.WriteFile(
		t,
		root,
		"pkg/format/delta.go",
		"package format\n\nimport \"fmt\"\n\nfunc Delta(kilo string) {\n\tfmt.Println(\"echo foxtrot golf hotel india\", \"juliett kilo lima mike november\", \"x\"+kilo)\n}\n",
	)
	o := option.New()
	o.Root = root
	o.Index = t.TempDir()
	printed := testutil.CaptureStdout(func() { gocheck.Print(checked(t, o)) })
	assert.StringContains(t, "pkg/format/delta.go:6: ", printed)
	assert.StringContains(t, " - on a line the fixes rewrite\n", printed)
}

func TestReportModeAnalysisNamesLineOnDisk(t *testing.T) {
	root := stackModule(t)
	testutil.WriteFile(
		t,
		root,
		"pkg/response/types.go",
		"package response\n\nfunc Alfa() int {\n\tbravo := 1\n\treturn bravo\n}\n\ntype LinkResponse struct {\n\tName string\n}\n\ntype Note struct {\n\tBody string\n}\n",
	)
	o := option.New()
	o.Root = root
	o.Index = t.TempDir()
	printed := testutil.CaptureStdout(func() { gocheck.Print(checked(t, o)) })
	assert.StringContains(
		t,
		"pkg/response/types.go:8: struct LinkResponse repeats the bag's name response - rename it Link\n",
		printed,
	)
}
