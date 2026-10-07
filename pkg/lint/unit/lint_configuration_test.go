package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/repository"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"strings"
	"testing"
)

func TestCheckReflowSkipsAConfiguredDirectory(t *testing.T) {
	long := join.Empty(strings.Repeat("word ", 17), "end\n")
	v := virtual_file_system.New()
	v.WriteString("pkg/runbook/Alert.md", long)
	v.WriteString("doc/prose.md", long)
	var r output.Results
	o := option.New("", false)
	o.ReflowSkips = []string{"pkg/runbook/"}
	fixes := lint.Check(repository.New(t.TempDir(), v), o, &r)
	assert.Boolean(t, false, fixes.Has("pkg/runbook/Alert.md"))
	assert.Boolean(t, true, fixes.Has("doc/prose.md"))
}

func TestCheckReflowRunsEverywhereWithoutSkips(t *testing.T) {
	long := join.Empty(strings.Repeat("word ", 17), "end\n")
	v := virtual_file_system.New()
	v.WriteString("pkg/runbook/Alert.md", long)
	var r output.Results
	fixes := lint.Check(
		repository.New(t.TempDir(), v),
		option.New("", false),
		&r,
	)
	assert.Boolean(t, true, fixes.Has("pkg/runbook/Alert.md"))
}
