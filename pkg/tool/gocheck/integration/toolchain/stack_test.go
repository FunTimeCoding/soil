package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gocheck"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"path/filepath"
	"testing"
)

func TestFixRefusalWithholdsEveryStagesWrites(t *testing.T) {
	root := stackModule(t)
	o := option.New()
	o.Root = root
	o.Fix = true
	o.Index = t.TempDir()
	o.BeforeCommit = func() {
		testutil.WriteFile(
			t,
			root,
			"pkg/spacing/alfa.go",
			"package spacing\n\nfunc Alfa() int {\n\treturn 2\n}\n",
		)
	}
	blocked := false
	printed := testutil.CaptureStdout(
		func() { blocked = gocheck.Print(checked(t, o)) },
	)
	assert.True(t, blocked)
	assert.String(
		t,
		"package format\n\nimport \"fmt\"\n\nfunc Delta() {\n\tfmt.Println(\"echo foxtrot golf hotel india\", \"juliett kilo lima mike november oscar\", \"x\")\n}\n",
		testutil.ReadFile(t, filepath.Join(root, "pkg/format/delta.go")),
	)
	assert.StringContains(
		t,
		"pkg/spacing/alfa.go: changed since it was read - nothing written, run again",
		printed,
	)
}

func TestFixCommitsEveryStageTogetherThroughEachOther(t *testing.T) {
	root := stackModule(t)
	testutil.WriteFile(
		t,
		root,
		"pkg/both/golf.go",
		"package both\n\nimport \"fmt\"\n\nfunc Golf() int {\n\thotel := 1\n\treturn hotel\n}\n\nfunc India() {\n\tjulietMsg := \"x\"\n\tfmt.Println(\"echo foxtrot golf hotel india\", \"juliett kilo lima mike november oscar\", julietMsg)\n}\n",
	)
	o := option.New()
	o.Root = root
	o.Fix = true
	o.Index = t.TempDir()
	o.BeforeCommit = func() {
		assert.String(
			t,
			"package spacing\n\nfunc Alfa() int {\n\tbravo := 1\n\treturn bravo\n}\n",
			testutil.ReadFile(t, filepath.Join(root, "pkg/spacing/alfa.go")),
		)
	}
	testutil.CaptureStdout(func() { gocheck.Print(checked(t, o)) })
	assert.String(
		t,
		"package both\n\nimport \"fmt\"\n\nfunc Golf() int {\n\thotel := 1\n\n\treturn hotel\n}\n\nfunc India() {\n\tjulietMessage := \"x\"\n\tfmt.Println(\n\t\t\"echo foxtrot golf hotel india\",\n\t\t\"juliett kilo lima mike november oscar\",\n\t\tjulietMessage,\n\t)\n}\n",
		testutil.ReadFile(t, filepath.Join(root, "pkg/both/golf.go")),
	)
	assert.String(
		t,
		"package spacing\n\nfunc Alfa() int {\n\tbravo := 1\n\n\treturn bravo\n}\n",
		testutil.ReadFile(t, filepath.Join(root, "pkg/spacing/alfa.go")),
	)
}

func TestReportModeShowsOneDiffForEveryStage(t *testing.T) {
	root := stackModule(t)
	o := option.New()
	o.Root = root
	o.Index = t.TempDir()
	printed := testutil.CaptureStdout(func() { gocheck.Print(checked(t, o)) })
	assert.StringContains(
		t,
		filepath.Join(root, "pkg/spacing/alfa.go"),
		printed,
	)
	assert.StringContains(
		t,
		filepath.Join(root, "pkg/format/delta.go"),
		printed,
	)
	assert.StringContains(t, "+\n", printed)
}
