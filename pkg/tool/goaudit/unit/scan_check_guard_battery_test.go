package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"testing"
)

func TestGuardTestMissingFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/model_context/new.go",
		"package model_context\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertConcern(t, s[0], constant.MissingGuardTestKey)
}

func TestGuardTestPresentClean(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/model_context/new.go",
		"package model_context\n",
	)
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	v.WriteString(
		"pkg/tool/gotestd/integration/guard/guard_test.go",
		"package guard\n",
	)
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertNoConcern(t, s[0], constant.MissingGuardTestKey)
}

func TestGuardTestUnguardedSkipped(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/worker/w.go", "package worker\n")
	v.WriteString("pkg/tool/gotestd/option/o.go", "package option\n")
	v.WriteString("pkg/tool/gotestd/run.go", "package gotestd\n")
	s := scan.Services(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(s))
	assertNoConcern(t, s[0], constant.MissingGuardTestKey)
}
