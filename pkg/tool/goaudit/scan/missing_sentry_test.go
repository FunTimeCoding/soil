package scan

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"testing"
)

func TestMissingSentryFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("cmd/gotest/main.go", "package main\n")
	v.WriteString(
		"pkg/tool/gotest/main.go",
		"package gotest\n\nfunc Main() {}\n",
	)
	result := MissingSentry(v)
	assert.Integer(t, 1, len(result))
	assert.String(t, MissingSentryKey, result[0].Key)
	assert.String(t, "cmd/gotest", result[0].Path)
}

func TestMissingSentryWithReporter(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("cmd/gotest/main.go", "package main\n")
	v.WriteString(
		"pkg/tool/gotest/run.go",
		"package gotest\n\nimport \"github.com/funtimecoding/soil/pkg/errors/sentry/reporter\"\n\nfunc Run() {\n\tr := reporter.New(\"gotest\")\n}\n",
	)
	result := MissingSentry(v)
	assert.Integer(t, 0, len(result))
}

func TestMissingSentryWithOptionalReporter(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("cmd/gotest/main.go", "package main\n")
	v.WriteString(
		"pkg/tool/gotest/run.go",
		"package gotest\n\nimport \"github.com/funtimecoding/soil/pkg/errors/sentry/reporter\"\n\nfunc Run() {\n\tr := reporter.NewOptional(\"gotest\", \"v0.1.0\")\n}\n",
	)
	result := MissingSentry(v)
	assert.Integer(t, 0, len(result))
}

func TestMissingSentrySkipsExample(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("cmd/example/main.go", "package main\n")
	result := MissingSentry(v)
	assert.Integer(t, 0, len(result))
}

func TestMissingSentrySorted(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("cmd/gobravo/main.go", "package main\n")
	v.WriteString("pkg/tool/gobravo/main.go", "package gobravo\n")
	v.WriteString("cmd/goalfa/main.go", "package main\n")
	v.WriteString("pkg/tool/goalfa/main.go", "package goalfa\n")
	result := MissingSentry(v)
	assert.Integer(t, 2, len(result))
	assert.String(t, "cmd/goalfa", result[0].Path)
	assert.String(t, "cmd/gobravo", result[1].Path)
}
