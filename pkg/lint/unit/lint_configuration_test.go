package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/repository"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckReflowSkipsConfiguredDirectory(t *testing.T) {
	long := join.Empty(strings.Repeat("word ", 17), "end\n")
	v := virtual_file_system.New()
	v.WriteString("pkg/runbook/Alert.md", long)
	v.WriteString("doc/prose.md", long)
	var r output.Results
	o := option.New("", false)
	o.ReflowSkips = []string{"pkg/runbook/"}
	root := t.TempDir()
	fixes := lint.Check(repository.New(root, v), o, &r)
	assert.Boolean(
		t,
		false,
		fixes.Holds(filepath.Join(root, "pkg/runbook/Alert.md")),
	)
	assert.Boolean(t, true, fixes.Holds(filepath.Join(root, "doc/prose.md")))
}

func TestCheckReflowRunsEverywhereWithoutSkips(t *testing.T) {
	long := join.Empty(strings.Repeat("word ", 17), "end\n")
	v := virtual_file_system.New()
	v.WriteString("pkg/runbook/Alert.md", long)
	var r output.Results
	root := t.TempDir()
	fixes := lint.Check(repository.New(root, v), option.New("", false), &r)
	assert.Boolean(
		t,
		true,
		fixes.Holds(filepath.Join(root, "pkg/runbook/Alert.md")),
	)
}
